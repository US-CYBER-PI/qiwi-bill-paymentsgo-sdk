package src

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpRequest(body []byte, url, method string, headers map[string]string) []byte {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	return responseBody
}
