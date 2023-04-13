package models

import "time"

type Users struct {
	ID        string     `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email,omitempty"`
	Role      string     `json:"role,omitempty""`
	CardID    string     `json:"card_id,omitempty"`
	SantriID  string     `json:"santri_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
