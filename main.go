package main

import (
	"fmt"
	comp "go_basic/cmd/10_composition"
)

func main() {
	notifiers := []comp.Notifier{
		&comp.Email{Address: "jay@choi.com"},
		&comp.SMS{Phone: "010-1234-5678"},
		&comp.Slack{Channel: "일반"},
	}

	results := comp.SendAll(notifiers, "공지사항")
	for _, r := range results {
		fmt.Println(r)
	}
}
