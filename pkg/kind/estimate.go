package kind

import "strings"

type Estimate struct {
	FirstName   string `json:"first_name,omitempty" yaml:"first-name"`
	LastName    string `json:"last_name,omitempty" yaml:"last-name"`
	Email       string `json:"email,omitempty" yaml:"email"`
	PhoneNumber string `json:"phone_number,omitempty" yaml:"phone-number"`
	Theme       string `json:"theme,omitempty" yaml:"theme"`
	Quantity    int    `json:"quantity,omitempty" yaml:"quantity"`
	NeededBy    string `json:"needed_by,omitempty" yaml:"needed-by"`
	Message     string `json:"message,omitempty" yaml:"message"`
}

func (e Estimate) HasContactInfo() bool {
	return strings.TrimSpace(e.FirstName) != "" &&
		strings.TrimSpace(e.LastName) != "" &&
		strings.TrimSpace(e.Email) != ""
}
