package composition

import "fmt"

type Email struct {
	Address string
}

func (e *Email) Notify(message string) string {
	return fmt.Sprintf("Email to %s: %s", e.Address, message)
}
