package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	//	pass := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6ImNoYWxsZW5nZSIsImV4cCI6MTc5MjkxODczOSwiaWF0IjoxNzkyOTE2OTM5LCJqdGkiOiI2MTM1YmY3M2FjNjVhNWUyMTUyMzA2YmMwMzM2ZTMwYiJ9.mIIZHkuWwBC3BjRl66-qnzH4LY4B56LME2NZ8YJVD6U"
	pass2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6InJld2FyZCIsInNpZCI6IjYxMzViZjczYWM2NWE1ZTIxNTIzMDZiYzAzMzZlMzBiIiwiZXhwIjoxNzkyOTE3NTU1LCJpYXQiOjE3OTI5MTY5NTUsImp0aSI6IjUyZTNiMzY2ZTMwMjIyZjBiZDcyOGUzNmFjODRkOTA0In0.fP_uBjAJ_aEIBGE-VBmP-XN0k4KuNOMhI4z5zipGJh4"
	req, _ := http.NewRequest("GET", "http://117.72.42.214:8080/api/reward", nil)
	req.Header.Set("Authorization", "Bearer "+pass2)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	resbody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(resbody))

	// seqs := []int{1,2,3,4,5,6,7,8,9,10}
	// for _, seq := range seqs {
	// 	go func(n int) {
	// 		body := strings.NewReader(fmt.Sprintf(`{"seq":%d}`, n))
	// 		req, _ := http.NewRequest("POST", "http://117.72.42.214:8080/api/challenge", body)
	// 		req.Header.Set("Authorization", "Bearer "+pass)
	// 		req.Header.Set("Content-Type", "application/json")
	// 		resp, _ := client.Do(req)
	// 		defer resp.Body.Close()
	// 	}(seq)
	// }

}

// 你已经找到鱼饼保险箱，但还需要凑齐 10 个编号才能打开。

// 请向当前地址 /api/challenge 发送 POST 请求：
// - 继续携带 Authorization: Bearer <你的通行证>
// - 设置 Content-Type: application/json
// - 请求体示例：{"seq": 1}

// 从第一个有效编号被收到开始，你有 5 秒时间提交 1～10。
// 每个请求只提交一个编号，编号不能重复计数。
// 阅读本说明不会启动计时。

// 提示：
// - 所有请求必须复用同一张通行证，不要每次重新领票。
// - 可以使用 go 关键字并发发送这 10 个请求。
// - 主程序需要等待请求完成，并读取每个请求的响应。
// - 集齐编号且顺序符合要求时，那次响应会给出新的通行证和奖励入口。
// - 可以先打乱编号再并发发送；并发并不保证服务端收到的顺序一定符合要求。
// - 超时或因完整升序、完整降序被拒绝后，下一次有效提交会自动开始新一轮。
// - 如果错过通关响应，再用原通行证 GET 本接口即可取回通关信息。

// 当前没有进行中的挑战。第一个有效提交将开始 5 秒计时。
