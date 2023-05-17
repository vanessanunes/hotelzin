package models

import "errors"

type Customer struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Document    string `json:"document"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
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
