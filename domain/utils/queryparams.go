package utils

type CustomerParams struct {
	Name     string `in:"query=name"`
	Phone    string `in:"query=phone"`
	Document string `in:"query=document"`
}
