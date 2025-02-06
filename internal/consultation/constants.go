package consultation

// Consultation States as string constants
const (
	// Initial States
	ConsultationCreated  = "created"
	ConsultationPending  = "pending"
	ConsultationAccepted = "accepted"
	ConsultationRejected = "rejected"
	ConsultationExpired  = "expired"

	// Call States
	ConsultationInProgress   = "in_progress"
	ConsultationCompleted    = "completed"
	ConsultationDisconnected = "disconnected"
	ConsultationForceStopped = "force_stopped"

	// End States
	ConsultationCancelled        = "cancelled"
	ConsultationNoShow           = "no_show"
	ConsultationBillingProcessed = "billing_processed"
	ConsultationRefundRequested  = "refund_requested"
	ConsultationRefundProcessed  = "refund_processed"
	ConsultationFeedbackGiven    = "feedback_given"
	ConsultationPenaltyApplied   = "penalty_applied"

	// Edge Cases
	ConsultationAttemptReconnect     = "attempt_reconnect"
	ConsultationAutoEndIfNoReconnect = "auto_end_if_no_reconnect"
	ConsultationPartialRefund        = "partial_refund"
	ConsultationRefundReviewed       = "refund_reviewed"
	ConsultationUserFlagged          = "user_flagged"
	ConsultationAstrologerFlagged    = "astrologer_flagged"
	ConsultationUserCanRetry         = "user_can_retry"

	// Consultation Type
	CallConsultation  = "call"
	ChatConsultation  = "chat"
	VideoConsultation = "video"

	// Activities
	SendNotificationToConsultantActivity = "SendNotificationToConsultantActivity"
)
