package miniorder

import (
	"fmt"
)

func ProcessOrder(inventory *Inventory, order *Order, payment Payment, receipt *ReceiptWriter) error {
	rollback := make(map[string]int, len(order.Items))

	for _, item := range order.Items {
		if err := inventory.RemoveStock(item.ProductID, item.Quantity); err != nil {
			for productID, rollbackAmount := range rollback {
				inventory.AddStock(productID, rollbackAmount)
			}

			return err
		}

		rollback[item.ProductID] += item.Quantity
	}

	if err := payment.Pay(order.Total()); err != nil {
		for productID, rollbackAmount := range rollback {
			inventory.AddStock(productID, rollbackAmount)
		}

		return &PaymentError{
			Method: fmt.Sprintf("%T", payment),
			Reason: err,
		}
	}

	order.Status = "completed"

	fmt.Fprintf(receipt, "%s", order.String())
	return nil
}
