package miniorder

import (
	"sync"
)

type Inventory struct {
	stock map[string]int
	mu    sync.Mutex
}

func NewInventory() *Inventory {
	return &Inventory{
		stock: map[string]int{},
	}
}

func (i *Inventory) AddStock(productID string, quantity int) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if err := validateQuantity(quantity); err != nil {
		return err
	}

	if _, ok := i.stock[productID]; !ok {
		return ErrProductNotFound
	}

	i.stock[productID] += quantity

	return nil
}

func (i *Inventory) GetProduct(productID string) (*Product, error) {
	if product, ok := i.stock[productID]; ok {
		return product, nil
	}

	return nil, ErrProductNotFound
}

func (i *Inventory) RemoveStock(productID string, quantity int) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if err := validateQuantity(quantity); err != nil {
		return err
	}

	if v, ok := i.stock[productID]; ok {
		if v >= quantity {
			i.stock[productID] -= quantity
			return nil
		}
		return ErrOutOfStock
	}
	return ErrProductNotFound
}

func (i *Inventory) HasStock(productID string, quantity int) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	if v, ok := i.stock[productID]; ok {
		return v >= quantity
	}

	return false
}

func validateQuantity(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}
