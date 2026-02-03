package main

import (
	"encoding/json"
	"fmt"
	setandtags "go_basic/cmd/15_set_and_tags"
)

func main() {
	a := &setandtags.Student{
		ID:      "S001",
		Name:    "",
		Classes: map[string]struct{}{},
	}
	b := &setandtags.Student{
		ID:      "S002",
		Name:    "이영희",
		Classes: map[string]struct{}{},
	}

	a.Enroll("Go기초")
	a.Enroll("알고리즘")
	a.Enroll("네트워크")

	b.Enroll("알고리즘")
	b.Enroll("데이터베이스")
	b.Enroll("네트워크")

	a.Enroll("임시과목")
	fmt.Println("Drop 전 - 임시과목 수강여부:", a.IsEnroll("임시과목"))
	a.Drop("임시과목")
	fmt.Println("Drop 후 - 임시과목 수강여부:", a.IsEnroll("임시과목"))

	fmt.Println("\n공통 과목:", setandtags.CommonCourses(a, b))
	fmt.Println("전체 과목:", setandtags.AllCourses(a, b))

	report := setandtags.BuildReport(a, b)
	jsonBytes, _ := json.MarshalIndent(report, "", "  ")
	fmt.Println("\n=== 리포트 ===")
	fmt.Println(string(jsonBytes))
}
