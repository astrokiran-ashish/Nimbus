package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"

	utils "github.com/astrokiran/nimbus/internal/common/utils"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	worflowutils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
)

func ConsultationHandShakeWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) error {

	logger := workflow.GetLogger(ctx)
	logger.Info("ConsultantStartWorkflow started", "state", state)

	// Send notification to consultant
	activityOptions := worflowutils.GetActivityOptions()
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	notification := map[string]string{
		"message":    "Consultation started",
		"event_type": "notification",
	}
	eventJsonString, err := utils.MapToJsonString(notification)
	if err != nil {
		return err
	}
	err = workflow.ExecuteActivity(ctx, constants.SendNotificationToConsultantActivity, state.ConsultantID, eventJsonString).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Add user to consultant waitlist
	notification = map[string]string{
		"user_id":         state.UserID,
		"user_name":       "test",
		"consultation_id": state.ConsultationID,
		"event_type":      "waitlist",
	}
	eventJsonString, err = utils.MapToJsonString(notification)
	if err != nil {
		return err
	}
	err = workflow.ExecuteActivity(ctx, constants.SendNotificationToConsultantActivity, state.ConsultantID, eventJsonString).Get(ctx, nil)
	if err != nil {
		return err
	}

	consultantResCh := workflow.GetSignalChannel(ctx, constants.SessionEndSignalCh)
	consultantTimer := workflow.NewTimer(ctx, 2*time.Minute)
	var consultantWaitlistAction workflowstates.ConsultantActionEvent

	// Flag to capture if the astrologer accepted.
	selector := workflow.NewSelector(ctx)
	var consultantAccepted bool

	selector.AddReceive(consultantResCh, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &consultantWaitlistAction)
		logger.Info("Recevied consultant response", "response", consultantWaitlistAction)
		if consultantWaitlistAction.Action == "accept" {
			consultantAccepted = true
		} else {
			consultantAccepted = false
		}
	})
	selector.AddFuture(consultantTimer, func(f workflow.Future) {
		logger.Info("Consultant did not respond in time")
		consultantAccepted = false
	})

	selector.Select(ctx)

	if !consultantAccepted {
		notification := map[string]string{
			"message":    "Astrologer Is Busy",
			"event_type": "notification",
		}
		eventJsonString, err := utils.MapToJsonString(notification)
		if err != nil {
			return err
		}
		err = workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil)
		if err != nil {
			return err
		}
		return nil
	}

	notification = map[string]string{
		"message":    "Astrologer accepted the call. You will get a call shortly",
		"event_type": "notification",
	}
	eventJsonString, err = utils.MapToJsonString(notification)
	if err != nil {
		return err
	}
	workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil)

	consultantOnCallCh := workflow.GetSignalChannel(ctx, constants.ConsultantOnCallCh)
	consultantTimerForInitiatingCall := workflow.NewTimer(ctx, 2*time.Minute)
	consultantInitiatedCall := false
	selector = workflow.NewSelector(ctx)

	selector.AddReceive(consultantOnCallCh, func(c workflow.ReceiveChannel, more bool) {
		logger.Info("Consultant is on call")
		notification = map[string]string{
			"message":    "Astrologer on Call. Please connect",
			"event_type": "call",
		}
		eventJsonString, err = utils.MapToJsonString(notification)
		workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil)
		consultantInitiatedCall = true
	})

	selector.AddFuture(consultantTimerForInitiatingCall, func(f workflow.Future) {
		logger.Info("Consultant did not response in time")
		notification = map[string]string{
			"message":    "Astrologer seems busy. Try and connect to other astrologer",
			"event_type": "call",
		}
		eventJsonString, err = utils.MapToJsonString(notification)
		workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil)
	})

	selector.Select(ctx)

	if !consultantInitiatedCall {
		logger.Info("Consultant did not initiate the call")
		return nil
	}

	var userAccepted bool
	var userResponse string

	userCh := workflow.GetSignalChannel(ctx, constants.UserResponseSignalCh)
	userTimer := workflow.NewTimer(ctx, 2*time.Minute)

	selector.AddFuture(userTimer, func(f workflow.Future) {
		logger.Info("User did not respond in time")
		userAccepted = false
	})

	selector.AddReceive(userCh, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &userResponse)
		logger.Info("Received user's response", "response", userResponse)
		userAccepted = true
	})

	if userAccepted {
		logger.Info("User accepted the call. Handshake complete")
		return nil
	}

	notification = map[string]string{
		"message":    "User did not respond in time",
		"event_type": "notification",
	}
	eventJsonString, err = utils.MapToJsonString(notification)
	if err != nil {
		return err
	}
	if err = workflow.ExecuteActivity(ctx, constants.SendNotificationToConsultantActivity, state.ConsultantID, eventJsonString).Get(ctx, nil); err != nil {
		return err
	}

	notification = map[string]string{
		"message":    "Session missed. Retry to connect",
		"event_type": "notification",
	}
	eventJsonString, err = utils.MapToJsonString(notification)
	if err != nil {
		return err
	}
	if err = workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil); err != nil {
		return err
	}

	UserWaitSignalCh := workflow.GetSignalChannel(ctx, constants.UserWaitSignalCh)
	var waitlistResponse string

	selector = workflow.NewSelector(ctx)
	selector.AddReceive(UserWaitSignalCh, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &waitlistResponse)
		logger.Info("Received user's response", "response", waitlistResponse)
	})
	selector.Select(ctx)

	if waitlistResponse == "yes" {
		notification = map[string]string{
			"message":    "You are moved to the top of the waitlist. Please wait for the astrologer to connect.",
			"event_type": "notification",
		}
		eventJsonString, err = utils.MapToJsonString(notification)
		if err != nil {
			return err
		}
		logger.Info("User chose to reconnect. Moving user to the top of the waitlist.")
		workflow.ExecuteActivity(ctx, constants.SendNotificationToUserActivity, state.UserID, eventJsonString).Get(ctx, nil)
	}

	logger.Info("User declined to reconnect. Ending handshake.")

	return nil
}
