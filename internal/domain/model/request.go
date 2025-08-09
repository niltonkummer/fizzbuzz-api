package model

type FizzBuzzRequest struct {
	Int1  int    `json:"int1" validate:"min=1"`
	Int2  int    `json:"int2" validate:"min=1"`
	Limit int    `json:"limit" validate:"min=0,max=500000"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

type FizzBuzzResponse struct {
	Response string `json:"response"`
}

type StatsResponse struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Hits  int    `json:"hits"`
}
