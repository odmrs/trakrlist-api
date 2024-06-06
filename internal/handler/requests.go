package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Genres   string `json:"genres"`
	Author   string `json:"author"`
	Duration uint   `json:"duration"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Name == "" && r.Type == "" && r.Genres == "" && r.Author == "" && r.Duration == 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Type == "" {
		return errParamIsRequired("type", "string")
	}
	if r.Genres == "" {
		return errParamIsRequired("genres", "string")
	}
	if r.Author == "" {
		return errParamIsRequired("author", "string")
	}
	if r.Duration == 0 {
		return errParamIsRequired("duration", "int")
	}

	return nil
}
