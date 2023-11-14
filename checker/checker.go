package checker

import (
	"Kraken/config"
	"Kraken/utils"
	"fmt"
	"h12.io/socks"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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

func CheckProxyWithClient(stats *utils.Stats, client *http.Client, Proxy string) {
	response, err := client.Get(config.GlobalConfig.CheckURL)
	if err != nil {
		utils.HandleError(err)
		stats.Invalid++
		stats.Checked++
		return
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.HandleError(err)
		stats.Errors++
		stats.Checked++
		return
	}
	if !strings.Contains(string(content), config.GlobalConfig.SuccessKey) {
		stats.Invalid++
		stats.Checked++
		return
	}

	stats.Valid++
	utils.AppendFile(config.GlobalConfig.OutputFolder+"/"+config.GlobalConfig.Prefix+".txt", Proxy)
}
