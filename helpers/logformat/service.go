package logformat

import (
	"encoding/json"
	"fmt"
)

// ServiceUnmarshallError default unmarshall error
func ServiceUnmarshallError(key string, err error) error {
	return fmt.Errorf("unable to unmarshal %v: %v", key, err)
}

// ServiceResponseCodeError default response code not 200 error
func ServiceResponseCodeError(status string) error {
	return fmt.Errorf("status not ok: %v", status)
}

// ServiceParseError default parse error
func ServiceParseError(key string, err error) error {
	return fmt.Errorf("unable to parse %v: %v", key, err)
}

// ServiceGetClientError default client get error
func ServiceGetClientError(err error) error {
	return fmt.Errorf("unable to get cient: %v", err)
}

// ServiceError default service error message
func ServiceError(msg string, values interface{}, err error) error {
	valuesString := ""
	if values != nil {
		jsonString, _ := json.Marshal(values)
		valuesString = string(jsonString)
	}
	return fmt.Errorf("service unable to %v %v: %v", msg, valuesString, err)
}
