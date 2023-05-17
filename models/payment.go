package models

type Payment struct {
	ID           int64   `json:"id"`
	BillID       int64   `json:"bill_id"`
	Value        float32 `json:"value"`
	TypePayment  string  `json:"type_payment"`
	Installments int     `json:"installments"`
}
