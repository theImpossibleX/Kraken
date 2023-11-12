package checker

import (
	"Kraken/config"
	"Kraken/utils"
	"encoding/json"
	"fmt"
	"h12.io/socks"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpResponse struct {
	CountryCode string `json:"countryCode"`
	Query       string `json:"query"`
}

func GetHttpTransport(Proxy string) *http.Transport {
	ProxyUrl, err := url.Parse(Proxy)
	if utils.HandleError(err) {
		return &http.Transport{}
	}

	return &http.Transport{
		Proxy: http.ProxyURL(ProxyUrl),
	}
}

func GetSocksTransport(Proxy string) *http.Transport {
	return &http.Transport{
		Dial: socks.Dial(fmt.Sprintf("%s?timeout=%ds", Proxy, time.Duration(config.GlobalConfig.Timeout)*time.Second)),
	}
}

func GetTransport(Proxy string) *http.Transport {
	if strings.Contains(Proxy, "http://") {
		return GetHttpTransport(Proxy)
	} else {
		return GetSocksTransport(Proxy)
	}
}

func ProxyReq(req string, proxy string) (res *http.Response, err error) {
	ReqUrl, err := url.Parse(req)
	if utils.HandleError(err) {
		return nil, err
	}
	client := &http.Client{
		Timeout:   time.Duration(config.GlobalConfig.Timeout) * time.Second,
		Transport: GetTransport(proxy),
	}
	res, err = client.Get(ReqUrl.String())
	return res, err
}

func CheckProxy(stats *utils.Stats, c, Proxy string) {
	response, err := ProxyReq(config.GlobalConfig.CheckURL, Proxy)
	if err != nil {
		stats.Invalid++
		stats.Checked++
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		stats.Errors++
		stats.Checked++
		return
	}

	if !strings.Contains(string(content), config.GlobalConfig.SuccessKey) {
		stats.Invalid++
		stats.Checked++
		return
	}

	var Resp HttpResponse
	err = json.Unmarshal(content, &Resp)
	if err != nil {
		stats.Invalid++
		stats.Checked++
		return
	}

	stats.Valid++
	utils.AppendFile(config.GlobalConfig.OutputFolder+"/"+config.GlobalConfig.Prefix+".txt", Proxy)
}
func CheckProxyWithClient(stats *utils.Stats, client *http.Client, Proxy string) {
	response, err := client.Get(config.GlobalConfig.CheckURL)
	if err != nil {
		stats.Invalid++
		stats.Checked++
		return
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		stats.Errors++
		stats.Checked++
		return
	}

	if !strings.Contains(string(content), config.GlobalConfig.SuccessKey) {
		stats.Invalid++
		stats.Checked++
		return
	}

	var Resp HttpResponse
	err = json.Unmarshal(content, &Resp)
	if err != nil {
		stats.Invalid++
		stats.Checked++
		return
	}

	stats.Valid++

	utils.AppendFile(config.GlobalConfig.OutputFolder+"/"+config.GlobalConfig.Prefix+".txt", Proxy)
}
