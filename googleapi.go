package structgraph

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiurl = "https://chart.googleapis.com/chart?cht=gv&chl="

func GenPngFromApi(str, filename string) (err error) {
	resp, err := http.Get(apiurl + url.QueryEscape(str))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(string(ret))
		return
	}
	err = ioutil.WriteFile(filename, ret, 0664)
	return
}
