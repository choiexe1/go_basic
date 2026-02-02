package composition

type Notifier interface {
	Notify(message string) string
}

func SendAll(notifiers []Notifier, message string) []string {
	result := []string{}

	for _, notifier := range notifiers {
		result = append(result, notifier.Notify(message))
	}

	return result
}
