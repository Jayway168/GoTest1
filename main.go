// Created by Jayway Tung on 2022-10-01
package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type CPUInfo struct {
	Sockets        int32  `json:"sockets"`
	CoresPerSocket int32  `json:"cores_per_socket"`
	ThreadsPerCore int32  `json:"threads_per_core"`
	CPU            int32  `json:"cpus"`
	CPUCores       int32  `json:"cpu_cores"`
	Processor      int32  `json:"processor"`
	Microcode      string `json:"microcode"`
	AddressSizes   string `json:"address_sizes"`
}

func init() {
	fmt.Println("Hello, Go world!")
}

func cpuInfo() {
	out, _ := exec.Command("lscpu").Output()
	outString := strings.TrimSpace(string(out))
	lines := strings.Split(outString, "\n")
	c := CPUInfo{}

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		fmt.Println("key = ", key)
		//fmt.Println("value = ", value)

		switch key {
		case "Socket(s)":
			fmt.Println("this is Socket(s)")
			t, _ := strconv.Atoi(value)
			c.Sockets = int32(t)
		case "Core(s) per socket":
			fmt.Println("this is Core(s) per socket")
			t, _ := strconv.Atoi(value)
			c.CoresPerSocket = int32(t)
		case "Thread(s) per core":
			fmt.Println("this is Thread(s) per core")
			t, _ := strconv.Atoi(value)
			c.ThreadsPerCore = int32(t)
		case "CPU(s)":
			fmt.Println("this is CPU(s)")
			t, _ := strconv.Atoi(value)
			c.CPU = int32(t)
		case "Address sizes":
			fmt.Println("this is Address sizes = ", value)
			c.AddressSizes = value
		case "CPUCores":
			fmt.Println("this is CPUCores")
			t, _ := strconv.Atoi(value)
			c.CPUCores = int32(t)
		case "Processor":
			fmt.Println("this is Processor")
			t, _ := strconv.Atoi(value)
			c.Processor = int32(t)
		case "Microcode":
			fmt.Println("this is Microcode")
			//t, _ := strings.TrimSpace(value)
			c.Microcode = value
		}
	}

	CPUInfoJSON, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(CPUInfoJSON))
}

func main() {

	var numCPU = runtime.NumCPU()
	fmt.Println("Number of threads = ", numCPU)

	//add a few lines of comments for push purpose
	// more
	// more
	// more
	// more
	// more
	cpuInfo()
}
