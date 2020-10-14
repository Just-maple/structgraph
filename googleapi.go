package structgraph

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	googleApiUrl     = "https://chart.googleapis.com/chart"
	quickChartApiUrl = "https://quickchart.io/graphviz"

	timeout = time.Second * 60
)

var cli = &http.Client{
	Timeout: timeout,
}

func GenPngFromQuickChartApi(str, filename string) (err error) {
	b, err := json.Marshal(map[string]string{
		"graph":  str,
		"format": "png",
	})
	if err != nil {
		return
	}
	resp, err := cli.Post(quickChartApiUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return
	}
	return write(resp, filename)
}

func write(resp *http.Response, filename string) (err error) {
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

func GenPngFromApi(str, filename string) (err error) {
	u := url.Values{}
	u.Add("cht", "gv")
	u.Add("chl", str)
	resp, err := cli.PostForm(googleApiUrl, u)
	if err != nil {
		return
	}
	return write(resp, filename)
}
