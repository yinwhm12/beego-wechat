package util

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func HttpGet(uri string)([]byte, error)  {
	response, err := http.Get(uri)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
