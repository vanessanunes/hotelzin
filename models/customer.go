package models

import "errors"

type Customer struct {
	ID          int64  `json:"customer_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Document    string `json:"document,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
}

type CustomerWithHosting struct {
	Customer
	Hostings   []Hosting `json:"bookings"`
	TotalValue float32   `json:"total_value"`
}

func (c *Customer) Validated() error {
	if c.Name == "" {
		return errors.New("Nome é obrigatório e não pode estar em branco")
	}
	if c.Document == "" {
		return errors.New("Documento é obrigatório e não pode estar em branco")
	}
	if c.PhoneNumber == "" {
		return errors.New("Número de telefone é obrigatório e não pode estar em branco")
	}
	return nil
}
