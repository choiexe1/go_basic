package channels

import (
	"fmt"
	"net/http"
)

func FetchOne(url string) string {
	// 채널 생성
	ch := make(chan string)

	// 고루틴에서..
	go func() {
		res, err := http.Get(url)
		// 요청 후에 채널에 메세지 전달
		// 여기서 끝나면, 바로 반환됨..
		// 채널에 보내기(ch <-)가 실행되면 대기 중인 받기(<-ch)가 즉시 통과됨.. 바로 리턴문에서 반환처리
		if err != nil {
			ch <- fmt.Sprintf("%s > error", url)
			return
		}

		// defer로 스트림 종료
		defer res.Body.Close()

		// 정상흐름에서, 채널에 메세지를 전달하기 때문에..
		// 여기서도 리턴문으로 가고 종료됨
		ch <- fmt.Sprintf("%s > %d", url, res.StatusCode)
	}()

	return <-ch
}
