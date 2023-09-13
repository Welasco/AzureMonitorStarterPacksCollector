package common

import (
	"fmt"
	"io"
	"net/http"
)

func Http_client(url string) (string, *http.Response, error) {
	//req, err := http.NewRequest("GET", "http://localhost/nginx_status", nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Fail to create HTTP request to NGINX")
		fmt.Println(err)
		return "", &http.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", &http.Response{}, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)

	return bodystr, resp, err

}
