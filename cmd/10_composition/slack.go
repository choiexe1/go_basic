package composition

import "fmt"

type Slack struct {
	Channel string
}

func (s *Slack) Notify(message string) string {
	return fmt.Sprintf("Slack #%s: %s", s.Channel, message)
}
