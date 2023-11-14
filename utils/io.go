package utils

import (
	"Kraken/config"
	"bufio"
	"fmt"
	"github.com/enriquebris/goconcurrentqueue"
	"os"
	"strings"
	"time"
)

func LoadProxies(filename string, prefix string) (*goconcurrentqueue.FIFO, error) {
	file, err := os.Open(filename)
	if err != nil {
		HandleError(err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			HandleError(err)
		}
	}(file)
	queue := goconcurrentqueue.NewFIFO()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := fmt.Sprint(scanner.Text())
		line = strings.ReplaceAll(line, "\r\n", "")
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, " ", "")
		line = strings.TrimSpace(line)
		err := queue.Enqueue(prefix + "://" + line)
		if err != nil {
			HandleError(err)
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return queue, nil
}
func AppendFile(FileName string, Content string) {
	File, err := os.OpenFile(fmt.Sprintf("%s", FileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if HandleError(err) {
		return
	}
	_, err = File.WriteString(Content + "\n")
	if HandleError(err) {
		return
	}
	File.Close()
}

func CreateFolderAndFiles() (folderName string, err error) {
	_ = os.Mkdir("Results", 0755)
	_ = os.Mkdir("data", 0755)
	folderName = time.Now().Format("02-01-2006 15-04-05")
	fullPath := "Results/" + folderName
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("error checking if directory exists: %v", err)
	}
	fileNames := []string{"http.txt", "socks5.txt", "socks4.txt"}
	for _, fileName := range fileNames {
		file, err := os.Create(fmt.Sprintf("%s/%s", fullPath, fileName))
		if err != nil {
			return "", fmt.Errorf("failed to create file %s: %v", fileName, err)
		}
		file.Close()
	}
	return fullPath, nil
}
func HandleError(Err error) bool {
	if Err != nil {
		if config.GlobalConfig.Debug {
			fmt.Println(Err)
			return true
		}
	}
	return false
}
