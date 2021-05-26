package main

import (
	"bytes"
	"easyhttp/component"
	"fmt"
	"time"
)

func main() {
	builder := component.NewClientBuilder()
	builder.SkipVerify(true)
	header := make(map[string]string)
	header["Content-Type"] = component.HTTP_CONTENT_TYPE_JSON
	builder.Header(header)
	client, err := builder.Build()
	if err != nil {
		panic("client creat error")
	}

	updateParams := `{"testKey":"123"}`
	//res := client.PostJson("http://127.0.0.1:8082/api/4", updateParams)
	//fmt.Println(res)

	client.PostJsonAsynWithParam("http://127.0.0.1:8082/api/4", updateParams, BadLabelOne)

	time.Sleep(time.Second * 10)

}

func say(call func(a, b int) int) int {
	return call(1, 2)
}

//印迹服务
func BadLabelOne(response component.IResponse, data interface{}) {
	fmt.Println(response.StatusCode())
	fmt.Println("data", bytes.NewBuffer(data.([]byte)))
}
