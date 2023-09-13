package common

import (
	"fmt"
	"io"
	"net/http"
)

func Http_client(url string) (string, error) {
	//req, err := http.NewRequest("GET", "http://localhost/nginx_status", nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)

	return bodystr, nil

}
