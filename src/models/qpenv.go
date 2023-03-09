package models

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	ENVIRONMENT         = "APP_ENV"
	TITLE               = "APP_TITLE"
	DEBUG_REQUESTS      = "DEBUGREQUESTS"
	DEBUG_JSON_MESSAGES = "DEBUGJSONMESSAGES"
)

type Environment struct{}

var ENV Environment

func (_ *Environment) IsDevelopment() bool {
	environment, _ := GetEnvStr(ENVIRONMENT)
	if strings.ToLower(environment) == "development" {
		return true
	} else {
		return false
	}
}

func (_ *Environment) AppTitle() string {
	title, _ := GetEnvStr(TITLE)
	return title
}

func (_ *Environment) DEBUGRequests() bool {

	if ENV.IsDevelopment() {
		environment, err := GetEnvBool(DEBUG_REQUESTS, true)
		if err == nil {
			return environment
		}
	}

	return false
}

func (_ *Environment) DEBUGJsonMessages() bool {

	if ENV.IsDevelopment() {
		environment, err := GetEnvBool(DEBUG_JSON_MESSAGES, true)
		if err == nil {
			return environment
		}
	}

	return false
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func GetEnvBool(key string, value bool) (bool, error) {
	result := value
	s, err := GetEnvStr(key)
	if err == nil {
		trying, err := strconv.ParseBool(s)
		if err == nil {
			result = trying
		}
	}
	return result, err
}

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}
