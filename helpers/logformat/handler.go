package logformat

import (
	"encoding/json"
	"fmt"
)

const (
	// InternalServerError default message
	InternalServerError = "internal server error"
	// BadRequest default message
	BadRequest = "bad request"
	// NotFound default message
	NotFound = "not found"
	// WebsiteInputValidationFailed defines a default message error for failed website input validation
	WebsiteInputValidationFailed = "website input validation failed"
	// InvalidCharacter default message
	InvalidCharacter = "invalid character used"
	// AuthenticationFailed default authentication failed message
	AuthenticationFailed = "Authentication failed"
)

// HandlerPostgresQueryError default postgres query error message
func HandlerPostgresQueryError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("unable to %v %v: %w [Postgres]", msg, valueString, err)
}

// HandlerPostgresQuerySuccess default postgres query error message
func HandlerPostgresQuerySuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("success %v %v [Postgres]", msg, valueString)
}

// HandlerInvalidParameter default invalid parameter error message
func HandlerInvalidParameter(key string) error {
	return fmt.Errorf("invalid %v parameter", key)
}

// HandlerInvalidRequest default request body decode error message
func HandlerInvalidRequest(err error) error {
	return fmt.Errorf("invalid request: %w", err)
}

// HandlerConvertValueError default handler convert value error message
func HandlerConvertValueError(key string, value interface{}, err error) error {
	return fmt.Errorf("unable to convert %v value %v: %w", key, value, err)
}

// HandlerValidateIPHostname default ip/hostname validator error message
func HandlerValidateIPHostname(value interface{}) error {
	return fmt.Errorf("invalid ip/hostname value %v", value)
}

// HandlerValidateURL default ip/hostname validator error message
func HandlerValidateURL(value interface{}) error {
	return fmt.Errorf("invalid url format %v", value)
}

// HandlerValidateWebsiteInput default ip/hostname validator error message
func HandlerValidateWebsiteInput(value interface{}) error {
	return fmt.Errorf("there are some missing value in website %v", value)
}

// HandlerCheckDataNotExist default request check data not exist error message
func HandlerCheckDataNotExist(key, value string) error {
	return fmt.Errorf("data %v value %v not exist", key, value)
}

// HandlerObjectInitializationError default object initialization error message
func HandlerObjectInitializationError(values string, err error) error {
	return fmt.Errorf("unable to initiate object %v: %w", values, err)
}

// HandlerNotFound default data not found error message
func HanderNotFound(key, value string) error {
	return fmt.Errorf("data %v value %v not found", key, value)
}

// HandlerEmptyValue default empty parameter error message
func HandlerEmptyValue(funcname, field string) error {
	return fmt.Errorf("%v invalid parameter: %v is empty", funcname, field)
}

// HandlerProducerError default Producer query error message
func HandlerProducerError(msg string, values interface{}, err error) error {
	jsonString, _ := json.Marshal(values)
	return fmt.Errorf("unable to %v %v: %v [Producer]", msg, string(jsonString), err)
}

// HandlerProducerSuccess default Producer query success message
func HandlerProducerSuccess(msg string, values interface{}) string {
	jsonString, _ := json.Marshal(values)
	return fmt.Sprintf("success %v %v [Producer]", msg, string(jsonString))
}

// HandlerParserError default parser error message
func HandlerParserError(msg string, values interface{}, err error) error {
	jsonString, _ := json.Marshal(values)
	return fmt.Errorf("unable to parse %v %v: %v [Parser]", msg, string(jsonString), err)
}

// HandlerConversionError default conversion error message
func HandlerConversionError(key, dataTYpe string, value interface{}) error {
	return fmt.Errorf("error type conversion %v: %v to %v", key, value, dataTYpe)
}

// HandlerError default handler error message
func HandlerError(msg string, values interface{}, err error) error {
	jsonString, _ := json.Marshal(values)
	return fmt.Errorf("unable to %v %v: %v [Handler]", msg, string(jsonString), err)
}

// HandlerDBQueryError default db query error message
func HandlerDBQueryError(msg string, values interface{}, err error) error {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Errorf("unable to %v %v: %w", msg, valueString, err)
}

// HandlerDBQuerySuccess default db query error message
func HandlerDBQuerySuccess(msg string, values interface{}) string {
	valueString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valueString = string(jsonString)
	}

	return fmt.Sprintf("Success %v %v", msg, valueString)
}

// HandlerUserAlreadyExist default user already exist message
func HandlerUserAlreadyExist(username interface{}) error {
	return fmt.Errorf("user %v is already exist", username)
}

// HandlerUserNotFound default user not found message
func HandlerUserNotFound(username interface{}) error {
	return fmt.Errorf("user %v not found", username)
}

// HandlerBadAuthorizationHeaderRequest default error message for bad authorization header resquest
func HandlerBadAuthorizationHeaderRequest() error {
	return fmt.Errorf("bad authorization header request")
}
