package setandtags

type Student struct {
	ID      string              `json:"id"`
	Name    string              `json:"name"`
	Classes map[string]struct{} `json:"classes"`
}

func (s *Student) Enroll(course string) {
	s.Classes[course] = struct{}{}
}

func (s *Student) Drop(course string) {
	delete(s.Classes, course)
}

func (s *Student) IsEnroll(course string) bool {
	_, ok := s.Classes[course]
	return ok
}

func CommonCourses(a, b *Student) []string {
	common := []string{}

	for course := range a.Classes {
		if b.IsEnroll(course) {
			common = append(common, course)
		}
	}

	return common
}

func AllCourses(a, b *Student) []string {
	all := map[string]struct{}{}

	for course := range a.Classes {
		all[course] = struct{}{}
	}

	for course := range b.Classes {
		all[course] = struct{}{}
	}

	result := []string{}
	for course := range all {
		result = append(result, course)
	}

	return result
}

func BuildReport(a, b *Student) any {
	allCourse := AllCourses(a, b)

	report := struct {
		AName       string   `json:"a_name,omitempty"`
		BName       string   `json:"b_name"`
		Common      []string `json:"common"`
		All         []string `json:"all"`
		TotalUnique int      `json:"total_unique"`
	}{
		AName:       a.Name,
		BName:       b.Name,
		Common:      CommonCourses(a, b),
		All:         allCourse,
		TotalUnique: len(allCourse),
	}

	return report
}
