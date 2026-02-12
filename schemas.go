package main

import (
	"fmt"
	"regexp"
	"time"
)

type CreatePersonRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type CreatePersonResponse struct {
	ID        int       `json:"id"`
	BookID    int       `json:"book_id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r CreatePersonRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	if r.Document == "" {
		return fmt.Errorf("document is required")
	}

	if len(r.Document) != 11 {
		return fmt.Errorf("document must have 11 numerics")
	}

	if !regexp.MustCompile(`^\d+$`).MatchString(r.Document) {
		return fmt.Errorf("document must have only numerics")
	}

	return nil
}
