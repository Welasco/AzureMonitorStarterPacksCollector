package common

import (
	"fmt"
	"io"
	"net/http"
)

// Http_client function
// This function is used to make HTTP requests to NGINX
// It returns the body of the response as string, the response itself and an error
func Http_client(url string) (string, *http.Response, error) {
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
