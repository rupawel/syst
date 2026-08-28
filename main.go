package main

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"os/exec"
	"os"
	"time"
	"fmt"
)

func cpuStat() {
	// infinite loop
	for {
		// clear the terminal
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// 1. CPU statistics
		cpuPercents, err := cpu.Percent(time.Second * 1, false)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("\n" + "								[STATISTICS]" + "\n")
			fmt.Printf("CPU:" + "\n")
			fmt.Printf("TOTAL CPU USAGE: %.2f%%" + "\n", cpuPercents[0])
		}
		cpuInfo, err := cpu.Info()
		if err == nil && len(cpuInfo) > 0 {
			fmt.Printf(" Model:	%s" + "\n", cpuInfo[0].ModelName)
			fmt.Printf(" Speed:	%.2f MHz" + "\n", cpuInfo[0].Mhz)
			fmt.Printf(" Cores:	%d" + "\n", len(cpuInfo))
		}

		time.Sleep(5*time.Second)
	}
}

func main() {
	cpuStat()
}
