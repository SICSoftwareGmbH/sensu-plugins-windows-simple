// Copyright © 2019 SIC! Software GmbH

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	warning := flag.Uint("w", 80, "warning value")
	critical := flag.Uint("c", 90, "critical value")

	flag.Parse()

	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get memory statistics: %s\n", err.Error())
		os.Exit(3)
	}

	infoMessage := fmt.Sprintf("system memory usage: %.2f%%", memory.UsedPercent)

	if memory.UsedPercent >= float64(*critical) {
		fmt.Printf("MEM CRITICAL: %s\n", infoMessage)
		os.Exit(2)
	} else if memory.UsedPercent >= float64(*warning) {
		fmt.Printf("MEM WARNING: %s\n", infoMessage)
		os.Exit(1)
	} else {
		fmt.Printf("MEM OK: %s\n", infoMessage)
		os.Exit(0)
	}
}
