// @title Go调用百度OCR接口测试
// @description
// @author wangpengliang
// @date 2022-03-30 15:51:24

package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// 通用文字识别高精度
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

	// 通用文字识别高精度含位置
	// var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"

	// 通用文字识别标准版
	// var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

	// 通用文字识别标准版含位置
	// var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"

	// 医疗费用明细
	// var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_detail"

	var accessToken = "24.a82b07c9d352ce36a11fbf585abdca92.2592000.1651816067.282335-25872618"

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile("C:/Users/Administrator/Desktop/小票样例/test.png")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendBody.Form.Add("language_type", "CHN_ENG")
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
