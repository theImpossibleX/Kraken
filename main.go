package main

import (
	"Kraken/checker"
	"Kraken/config"
	"Kraken/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var stats *utils.Stats

func init() {
	config.Load("config.json")
	log.SetOutput(ioutil.Discard)
	utils.PrintLogo()
	stats = utils.InitStats()
}

func StartChecker() {
	fmt.Println("Loading proxies")
	proxyQueue, _ := utils.LoadProxies(config.GlobalConfig.ProxyFilepath, config.GlobalConfig.Prefix)
	stats.Total = proxyQueue.GetLen()
	var clientPool []*http.Client
	for i := 0; i < config.GlobalConfig.Threads; i++ {
		proxy, err := proxyQueue.Dequeue()
		if err != nil {
			break
		}
		transport := checker.GetTransport(proxy.(string))
		clientPool = append(clientPool, &http.Client{Transport: transport})
	}

	go stats.ConsoleStats()
	go stats.CalcStats()

	var wg sync.WaitGroup
	for _, client := range clientPool {
		wg.Add(1)
		go func(c *http.Client) {
			defer wg.Done()
			for {
				proxy, err := proxyQueue.Dequeue()
				if err != nil {
					return
				}
				checker.CheckProxyWithClient(stats, c, proxy.(string))
			}
		}(client)
	}
	wg.Wait()
}

func main() {
	StartChecker()
}
