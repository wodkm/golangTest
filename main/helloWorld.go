package main

import (
	"fmt"
	"runtime"
)

func cpuNo() int {
	fmt.Println(runtime.NumCPU())
	return runtime.NumCPU()
}

func main() {
	fmt.Println("hello world")
	cpuNo()
	go cpuNo()
}
