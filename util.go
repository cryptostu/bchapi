package bchapi

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (str string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	str = string(body)
	return

}
