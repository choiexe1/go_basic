package string

import (
	"strings"
)

func Counter(values string) map[string]int {
	counts := map[string]int{}

	values = strings.ToLower(values)
	r := strings.SplitSeq(values, " ")

	for v := range r {
		counts[string(v)]++
	}

	return counts
}
