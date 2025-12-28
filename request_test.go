package gresty

import (
	"log"
	"testing"
	"time"
)

// TestRequest test request.
func TestRequest(t *testing.T) {
	s := New(WithTimeout(3 * time.Second))

	// 请求参数设置
	opt := &RequestOptions{
		// Params: map[string]interface{}{
		// 	"objid":   12784,
		// 	"objtype": 1,
		// 	"p":       0,
		// },
		RetryCount:         2,
		InsecureSkipVerify: true,
	}

	res := s.Do("get", "http://localhost:50051/healthz", opt)
	if res.Err != nil {
		log.Println("err: ", res.Err)
		return
	}

	log.Println("data: ", string(res.Body))

	// data := &ApiStdRes{}
	// err := res.Json(data)
	// log.Println(err)
	// log.Println(data.Code, data.Message)
	// log.Println(data.Data)
	//
	// res = s.Do("post", "http://localhost:1338/v1/data", &RequestOptions{
	// 	Data: map[string]interface{}{
	// 		"id": "1234",
	// 	},
	// 	RetryCount: 2, // 重试次数
	// })
	// if res.Err != nil {
	// 	log.Println("err: ", res.Err)
	// 	return
	// }
	//
	// log.Println(res.Err, string(res.Body))
}
