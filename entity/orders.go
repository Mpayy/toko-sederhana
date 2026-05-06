package entity

import "time"

type Orders struct {
	Id         int32
	ProductId  int32
	Quantity   int32
	TotalPrice int32
	CreatedAt  time.Time
}

type OrderResponse struct {
	Id          int32
	ProductName string
	Quantity    int32
	TotalPrice  int32
	CreatedAt   time.Time
}
