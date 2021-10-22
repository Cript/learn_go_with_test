package main

import (
	"main/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}