package miniorder

import (
	"fmt"
	"time"
)

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Subtotal  int    `json:"subtotal"`
}

type Order struct {
	ID        string      `json:"id"`
	Items     []OrderItem `json:"items"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func (o *Order) Total() int {
	var total int

	for _, v := range o.Items {
		total += v.Subtotal
	}

	return total
}

func (o *Order) String() string {
	return fmt.Sprintf("주문 %s | %d건 | 총 %d원 | %s", o.ID, len(o.Items), o.Total(), o.Status)
}
