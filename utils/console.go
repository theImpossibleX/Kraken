package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"strconv"
	"strings"
	"sync"
	"time"
)

var StartTime = time.Now()

type Stats struct {
	mu         sync.Mutex
	Valid      int
	Invalid    int
	Checked    int
	Total      int
	Errors     int
	Time       float64
	TTC        string
	CPM        int
	hasPrinted bool
}

func InitStats() *Stats {
	return &Stats{}
}

func (s *Stats) CalcStats() {
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()
	var lastCheckCount int
	for {
		select {
		case <-ticker.C:
			s.Time = time.Since(StartTime).Seconds()
			currentCheckCount := s.Checked
			if lastCheckCount != 0 {
				checksSinceLastInterval := float64(currentCheckCount - lastCheckCount)
				cpm := checksSinceLastInterval * (60.0 / 4.0)
				s.TTC = strconv.Itoa(int(float64(s.Total-currentCheckCount) / cpm))
				s.CPM = int(cpm)
			}
			lastCheckCount = currentCheckCount
		}
	}
}

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()
	color.Red(`
   __ __         __          
  / //_/______ _/ /_____ ___ 
 / ,< / __/ _ '/  '_/ -_) _ \
/_/|_/_/  \_,_/_/\_\\__/_//_/

Advanced Proxy Checker
`)
}

func (s *Stats) ConsoleStats() {
	for {
		s.mu.Lock()
		screen.Clear()
		PrintLogo()
		stats := []string{
			fmt.Sprintf("Valid   : %v", s.Valid),
			fmt.Sprintf("Invalid : %v", s.Invalid),
			fmt.Sprintf("Checked : %v", s.Checked),
			fmt.Sprintf("Total   : %v", s.Total),
			fmt.Sprintf("Errors  : %v", s.Errors),
			fmt.Sprintf("Time    : %v", s.Time),
			fmt.Sprintf("TTC     : %v", s.TTC),
			fmt.Sprintf("CPM     : %v", s.CPM),
		}

		maxLen := 0
		for _, stat := range stats {
			if len(stat) > maxLen {
				maxLen = len(stat)
			}
		}
		border := "+" + strings.Repeat("-", maxLen+2) + "+"
		//numLinesUp := len(stats) + 2
		//if s.hasPrinted {
		//	fmt.Printf("\033[%dA", numLinesUp)
		//}
		fmt.Println(border)
		for _, stat := range stats {
			fmt.Printf("| %-*s |\n", maxLen, stat)
		}
		fmt.Println(border)
		s.hasPrinted = true
		s.mu.Unlock()
		time.Sleep(time.Millisecond * 1000)

	}
}

func HandleError(Err error) bool {
	if Err != nil {
		//fmt.Println(Err)
		return true
	}
	return false
}
