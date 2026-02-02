package goroutines

import (
	"fmt"
	"net/http"
	"sync"
)

func DoA() string {
	return "A"
}

func DoB() string {
	return "B"
}

func DoC() string {
	return "C"
}

func FetchAll(urls []string) []string {
	var wg sync.WaitGroup

	// 레이스 컨디션이 발생하지 않게, 각 인덱스에 쓰기 작업
	results := make([]string, len(urls))

	// 반복문 돌면서..
	for i, url := range urls {
		// wait group에 카운트 추가 하고..
		wg.Add(1)
		go func(idx int, u string) {
			// defer로 함수 종료시 wait group 카운트 -1
			defer wg.Done()

			resp, err := http.Get(u)
			// 예외 있으면 슬라이스에 추가
			if err != nil {
				results[idx] = fmt.Sprintf("%s > error: %s", u, err)
				return
			}

			// defer로 함수 종료시 스트림 정리..
			defer resp.Body.Close()

			// 정상 응답이면, 상태코드랑 함께 슬라이스에 추가
			results[idx] = fmt.Sprintf("%s > %d", u, resp.StatusCode)

		}(i, url)
	}

	// wait group 작업 종료 기다리게..
	wg.Wait()

	// wait group 작업 끝나면 결과반환
	return results
}
