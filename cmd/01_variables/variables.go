package variables

import "fmt"

func Calculate() {
	var value float64
	var convert int

	fmt.Print("[SYSTEM] 값을 입력하세요.\n")
	fmt.Scan(&value)
	fmt.Print("[SYSTEM] 변환 방향을 설정하세요.\n")
	fmt.Print("[SYSTEM] 1: KM -> Mile, 2: Mile -> KM\n")
	fmt.Scan(&convert)

	switch convert {
	case 1:
		fmt.Printf("[SYSTEM] %.2f KM는 %.2f Mile입니다.\n", value, value*0.621371)
	case 2:
		fmt.Printf("[SYSTEM] %.2f Mile는 %.2f KM입니다.\n", value, value*1.60934)
	default:
		fmt.Println("[SYSTEM] 잘못된 입력입니다.")
	}
}
