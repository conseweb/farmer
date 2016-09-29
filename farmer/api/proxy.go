package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetClient() *http.Client {
	if proxyClient != nil {
		return proxyClient
	}
	return http.DefaultClient
}

func ProxyFabric(rw http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	cli := GetClient()

	req.RequestURI = ""

	req.URL.Host = daemon.GetRESTAddr()
	req.URL.Scheme = "http"
	if API_PREFIX != "" {
		req.URL.Path = req.URL.Path[len(API_PREFIX):]
	}
	req.Close = true
	req.Header.Set("Connection", "close")

	resp, err := cli.Do(req)
	if err != nil {
		ctx.Error(resp.StatusCode, err)
		return
	}

	for k, vv := range resp.Header {
		if strings.ToLower(k) == "content-length" {
			continue
		}
		for _, v := range vv {
			rw.Header().Set(k, v)
		}
	}

	rw.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		ctx.Error(resp.StatusCode, fmt.Errorf("%s", msg))
	} else {
		io.Copy(rw, resp.Body)
	}
}
