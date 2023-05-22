package models

import "errors"

type Checking struct {
	ID               *int64  `json:"id,omitempty" swaggerignore:"true"`
	BookingId        *int64  `json:"booking_id"`
	CheckingDatetime *string `json:"checking_datetime,omitempty" example:"2023-05-20 20:00:00"`
}

type Checkout struct {
	CheckoutDatetime *string `json:"checkout_datetime,omitempty" example:"2023-05-20 20:00:00"`
}

type CheckingComplete struct {
	Checking
	Checkout
	Status *string `json:"status,omitempty"`
}

func (b *Checking) Validated() error {
	if b.BookingId == nil {
		return errors.New("Por favor, selecione a reserva para realizar o pagamento")
	}
	if b.CheckingDatetime == nil {
		return errors.New("Para que a reserva seja validade, inclua o data e hora do checking")
	}
	return nil
}
