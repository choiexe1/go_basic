package miniorder

import (
	"fmt"
)

type Category int

const (
	Electronics Category = iota
	Food
	Clothing
)

func (c Category) String() string {
	switch c {
	case Electronics:
		return "Electronics"
	case Food:
		return "Food"
	case Clothing:
		return "Clothing"
	default:
		return "Unknown"
	}
}

func (c Category) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

type Product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Category Category `json:"category"`
}

func (p *Product) String() string {
	return fmt.Sprintf("[%s] %s (%d원)", p.Category.String(), p.Name, p.Price)
}
