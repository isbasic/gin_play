package main

import (
	"fmt"
	"time"

	_ "github.com/isbasic/gin_play/sample"
)

func setPort(i int) string {
	return fmt.Sprintf(":%d", i)
}

func main() {
	times := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			now := time.Now()
			writeTimes(now, times)
		}
		close(times)
	}()

	for t := range times {
		fmt.Println(t)
	}
}

func writeTimes(t time.Time, times chan<- string) {
	times <- fmtTime(t)
}

func fmtTime(t time.Time) string {
	resTmpl := "%04d%02d%02d%02d%02d%02d%09d" // yyyyMMddhhmmssnanosecond
	y, m, d := t.Date()
	h, mn, s, nano := t.Hour(), t.Minute(), t.Second(), t.Nanosecond()
	return fmt.Sprintf(resTmpl, y, m, d, h, mn, s, nano)
}
