package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type RespData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	res, err := req.Get(fmt.Sprintf("http://127.0.0.1:8000/add?a=%d&b=%d", a, b))

	if err != nil {
		fmt.Println(err)
	}
	body, _ := res.Body()

	rspData := RespData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}
func main() {
	fmt.Println(Add(1, 2))

}
