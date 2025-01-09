package configs

import (
	"encoding/json"
	"os"
	"strconv"
)

var config map[string]interface{}

func init() {
	deploymentType := os.Getenv("DEVELOPMENT_ENV")
	if deploymentType == "" {
		deploymentType = "stage" // default to stage if not set
	}

	configFile, err := os.Open(deploymentType + "/config.json")
	if err != nil {
		config = make(map[string]interface{})
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		config = make(map[string]interface{})
	}
}

func GetString(key, defaultValue string) string {
	if value, exists := config[key]; exists {
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) int {
	if value, exists := config[key]; exists {
		if floatValue, ok := value.(float64); ok {
			return int(floatValue)
		}
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func GetBool(key string, defaultValue bool) bool {
	if value, exists := config[key]; exists {
		if boolValue, ok := value.(bool); ok {
			return boolValue
		}
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}
