package main

import (
	"errors"
	"fmt"
	"os"
	"puzzlers/article37/common"
	"puzzlers/article37/common/op"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	if err := startCPUProfile(f); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}
	if err = common.Execute(op.CPUProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
