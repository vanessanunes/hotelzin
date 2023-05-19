package models

import "errors"

type Payment struct {
	ID           int64   `json:"id"`
	BillID       int64   `json:"bill_id"`
	Value        float32 `json:"value"`
	TypePayment  string  `json:"type_payment"`
	Installments int     `json:"installments"`
}

func (b *Payment) Validated() error {
	if b.BillID == 0 {
		return errors.New("Para gerar pagamento, por favor, inclua o ID da conta a ser paga")
	}
	if b.Value == 0 {
		return errors.New("O valor da conta é invalido. Por favor, acionar o gerente!")
	}
	if b.TypePayment == "" {
		b.TypePayment = "cash"
	}
	return nil
}
