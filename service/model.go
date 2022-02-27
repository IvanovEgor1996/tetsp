package service

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Money int

func (m *Money) UnmarshalJSON(bytes []byte) error {
	var f float64
	err := json.Unmarshal(bytes, &f)
	if err != nil {
		return err
	}
	*m = Money(f * 100)
	return nil
}

type IncRequest struct {
	User uuid.UUID `json:"user"`
	Sum  Money     `json:"sum"`
}

type MoveRequest struct {
	UserFrom uuid.UUID `json:"user_from"`
	UserTo   uuid.UUID `json:"user_to"`
	Sum      Money     `json:"sum"`
}
