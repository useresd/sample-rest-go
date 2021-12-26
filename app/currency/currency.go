package currency

import (
	"errors"
)

// Currency ...
type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}

// CreateDTO ...
type EditDTO struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}

// Validate ...
func (c *EditDTO) Validate() error {

	// Ensure the name is not empty
	if c.Name == "" {
		return errors.New("name is required")
	}

	// Ensure the abbreviation is not empty
	if c.Abbr == "" {
		return errors.New("abbr is required")
	}

	return nil
}
