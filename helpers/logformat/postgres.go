package logformat

import "fmt"

// PostgresQueryResponseError default postgres query get response error message
func PostgresQueryResponseError(err error) error {
	return fmt.Errorf("postgres query response error: %w", err)
}

// PostgresScanError default error postgres scan response error message
func PostgresScanError(err error) error {
	return fmt.Errorf("postgres scan error: %w", err)
}

// PostgresUnmarshalError default postgres unmarshal response error message
func PostgresUnmarshalError(err error) error {
	return fmt.Errorf("postgress unmarshal error: %w", err)
}

// PostgresExecResponseError default postgres exec get response error message
func PostgresExecResponseError(err error) error {
	return fmt.Errorf("postgres exec response error: %w", err)
}

// PostgresNewResponseError default postgres new get response error message
func PostgresNewResponseError(err error) error {
	return fmt.Errorf("postgres new response error: %w", err)
}

// PostgresError default postgres response error message
func PostgresError(err error) error {
	return fmt.Errorf("postgres error: %w", err)
}
