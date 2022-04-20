package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID string `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	ID         string    `json:"id" :"id"`
	Item       *OrderItem `json:"item"`
	Desc       string    `json:"desc,omitempty"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
}

func marshal()  {
	o := Order{
		ID:         "1234",
		Item: &OrderItem{
			ID: "item_1",
			Name: "learn go",
			Price: 10,
		},
		Quantity:   3,
		TotalPrice: 30,
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
func unmarshal()  {
	s:=`{"id":"1234","item":{"id":"item_1","name":"learn go","price":10},"quantity":3,"total_price":30}`
	var o Order
	err:=json.Unmarshal([]byte(s),&o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v",o)
}
func main() {
	unmarshal()
}
