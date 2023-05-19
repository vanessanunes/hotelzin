package models

import "errors"

type Checking struct {
	ID               int64  `json:"id"`
	BookingId        int64  `json:"booking_id"`
	CheckingDatetime string `json:"checking_datetime,omitempty"`
	CheckoutDatetime string `json:"checkout_datetime,omitempty"`
}

func (b *Checking) Validated() error {
	if b.BookingId == 0 {
		return errors.New("Por favor, selecione a reserva para realizar o pagamento")
	}
	if b.CheckingDatetime == "" {
		return errors.New("Para que a reserva seja validade, inclua o data e hora do checking")
	}
	return nil
}
