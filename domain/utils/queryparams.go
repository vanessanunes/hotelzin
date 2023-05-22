package utils

import "github.com/gorilla/schema"

var Decoder = schema.NewDecoder()

type CustomerParams struct {
	Name     string `in:"query=name"`
	Phone    string `in:"query=phone"`
	Document string `in:"query=document"`
}

type BookingParams struct {
	Status string `in:"query=status"`
}
