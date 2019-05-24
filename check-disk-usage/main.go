// Copyright © 2019 SIC! Software GmbH

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

var (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
)

func main() {
	ignoreTypesFlag := flag.String("x", "", "ignore disk types")
	warning := flag.Uint("w", 80, "usage warning value")
	critical := flag.Uint("c", 90, "usage critical value")
	warningInodes := flag.Uint("W", 80, "inodes warning value")
	criticalIndoes := flag.Uint("K", 90, "inodes critical value")

	flag.Parse()

	ignoreTypes := strings.Split(*ignoreTypesFlag, ",")

	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("Failed to get disk statistics: %s\n", err.Error())
		os.Exit(3)
	}

	status := OK
	infoMessage := ""

	for _, partition := range partitions {
		if partition.Fstype == "" {
			continue
		}

		if stringInList(partition.Fstype, ignoreTypes) {
			continue
		}

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Failed to get disk statistics for %s: %s\n", partition.Mountpoint, err.Error())
			os.Exit(3)
		}

		if usage.UsedPercent >= float64(*critical) {
			status = CRITICAL
		} else if usage.UsedPercent >= float64(*warning) {
			status = WARNING
		}

		if usage.InodesUsedPercent >= float64(*criticalIndoes) {
			status = CRITICAL
		} else if usage.InodesUsedPercent >= float64(*warningInodes) {
			status = WARNING
		}

		infoMessage = fmt.Sprintf("%s %s (usage=%.2f%% inodes=%.2f%%)", infoMessage, partition.Mountpoint, usage.UsedPercent, usage.InodesUsedPercent)
	}

	if status == CRITICAL {
		fmt.Printf("CRITICAL:%s\n", infoMessage)
		os.Exit(2)
	} else if status == WARNING {
		fmt.Printf("WARNING:%s\n", infoMessage)
		os.Exit(1)
	} else {
		fmt.Printf("OK:%s\n", infoMessage)
		os.Exit(0)
	}
}

func stringInList(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}
