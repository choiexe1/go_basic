package composition

import "fmt"

type SMS struct {
	Phone string
}

func (s *SMS) Notify(message string) string {
	return fmt.Sprintf("SMS to %s: %s", s.Phone, message)
}
