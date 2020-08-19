package types

import "github.com/google/uuid"

type (
	//Boolean : The main struct, used in the DB
	Boolean struct {
		ID    uuid.UUID `json:"id"`
		Value bool      `json:"value"`
		Key   string    `json:"key"`
	}

	//InputBoolean : Struct received in the request body from user
	InputBoolean struct {
		Value bool   `json:"value"`
		Key   string `json:"key"`
	}
)
