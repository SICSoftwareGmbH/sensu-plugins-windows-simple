// Copyright © 2019 SIC! Software GmbH

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	warning := flag.Uint("w", 80, "warning value")
	critical := flag.Uint("c", 100, "critical value")

	flag.Parse()

	times, err := cpu.Times(false)
	if err != nil {
		fmt.Printf("Failed to get CPU statistics: %s\n", err.Error())
		os.Exit(3)
	}
	if len(times) != 1 {
		fmt.Printf("Failed to get CPU statistics\n")
		os.Exit(3)
	}

	infoBefore := times[0]

	time.Sleep(5 * time.Second)

	times, err = cpu.Times(false)
	if err != nil {
		fmt.Printf("Failed to get CPU statistics: %s\n", err.Error())
		os.Exit(3)
	}
	if len(times) != 1 {
		fmt.Printf("Failed to get CPU statistics\n")
		os.Exit(3)
	}

	infoNow := times[0]

	diffUser := infoNow.User - infoBefore.User
	diffNice := infoNow.Nice - infoBefore.Nice
	diffSystem := infoNow.System - infoBefore.System
	diffIdle := infoNow.Idle - infoBefore.Idle
	diffIowait := infoNow.Iowait - infoBefore.Iowait
	diffIrq := infoNow.Irq - infoBefore.Irq
	diffSoftirq := infoNow.Softirq - infoBefore.Softirq
	diffSteal := infoNow.Steal - infoBefore.Steal
	diffGuest := infoNow.Guest - infoBefore.Guest
	diffGuestNice := infoNow.GuestNice - infoBefore.GuestNice

	totalDiff := diffUser + diffNice + diffSystem + diffIdle + diffIowait + diffIrq + diffSoftirq + diffSteal + diffGuest + diffGuestNice
	idleDiff := diffIdle + diffIowait + diffSteal + diffGuest + diffGuestNice
	usage := 100 * (totalDiff - idleDiff) / totalDiff

	infoMessage := fmt.Sprintf(
		"total=%.2f user=%.2f nice=%.2f system=%.2f idle=%.2f iowait=%.2f irq=%.2f softirq=%.2f steal=%.2f guest=%.2f guest_nice=%.2f",
		usage,
		(100 * (diffUser / totalDiff)),
		(100 * (diffNice / totalDiff)),
		(100 * (diffSystem / totalDiff)),
		(100 * (diffIdle / totalDiff)),
		(100 * (diffIowait / totalDiff)),
		(100 * (diffIrq / totalDiff)),
		(100 * (diffSoftirq / totalDiff)),
		(100 * (diffSteal / totalDiff)),
		(100 * (diffGuest / totalDiff)),
		(100 * (diffGuestNice / totalDiff)),
	)

	if usage >= float64(*critical) {
		fmt.Printf("TOTAL CRITICAL: %s\n", infoMessage)
		os.Exit(2)
	} else if usage >= float64(*warning) {
		fmt.Printf("TOTAL WARNING: %s\n", infoMessage)
		os.Exit(1)
	} else {
		fmt.Printf("TOTAL OK: %s\n", infoMessage)
		os.Exit(0)
	}
}
