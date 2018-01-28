package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

const (
	STAT_DISCHARGING = "Discharging"
	DELAY            = 8 * time.Second
	CMD_SHUTOFF      = "s2ram"
)

func main() {
	discharging := false
	for true {
		data, err := ioutil.ReadFile("/sys/class/power_supply/BAT1/status")
		if err != nil {
			panic(err)
		}
		status := strings.TrimSpace(string(data))
		if status == STAT_DISCHARGING {
			if discharging {
				discharging = false
				shutoff()
			} else {
				discharging = true
			}
		} else {
			discharging = false
		}
		time.Sleep(DELAY)
	}
}

func shutoff() {
	exec.Command(CMD_SHUTOFF).Start()
	time.Sleep(1 * time.Second)
}
