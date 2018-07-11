package model

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Service02Model struct {
	Code  int    `json:"code"`
	Items []Item `json:"items"`
}
