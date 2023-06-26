package generate

import "github.com/google/uuid"

// CreateUUID is function to generate uuid
func CreateUUID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
