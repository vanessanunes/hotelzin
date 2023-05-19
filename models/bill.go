package models

import "errors"

type Bill struct {
	ID         int64   `json:"id"`
	BookingId  int64   `json:"booking_id"`
	ExtraHour  int32   `json:"extra_hour"`
	TotalValue float32 `json:"total_value"`
}

func (b *Bill) Validated() error {
	if b.BookingId == 0 {
		return errors.New("Por favor, selecione a reserva para realizar o pagamento")
	}
	if b.TotalValue == 0 {
		return errors.New("Valor de conta incorreto, por favor, chamar o gerente!")
	}
	return nil
}
