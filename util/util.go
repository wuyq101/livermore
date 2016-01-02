package util

import (
	"github.com/wuyq101/livermore/logger"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		logger.Error("Failed to fetch url %s, err %+v", url, err)
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.Error("Failed to read data for url %s, err %+v", url, err)
		return nil, err
	}
	return result, err
}
