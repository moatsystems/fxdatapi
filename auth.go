// auth.go
package fxdatapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Authenticate(username, password string) string {
	url := "https://fxdatapi.com/v1/login"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
  "username": "%s",
  "password": "%s"
}`, username, password))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
