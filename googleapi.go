package structgraph

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	apiUrl = "https://chart.googleapis.com/chart"

	timeout = time.Second * 60
)

var cli = &http.Client{
	Timeout: timeout,
}

func GenPngFromApi(str, filename string) (err error) {
	u := url.Values{}
	u.Add("cht", "gv")
	u.Add("chl", str)
	resp, err := cli.PostForm(apiUrl, u)
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
