package structgraph

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiurl = "https://chart.googleapis.com/chart"

func GenPngFromApi(str, filename string) (err error) {
	u := url.Values{}
	u.Add("cht", "gv")
	u.Add("chl", str)
	resp, err := http.PostForm(apiurl, u)
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
