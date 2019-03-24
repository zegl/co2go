package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	gramsOfco2perkwh = 64 // Average in France
	cpuThreads       = 54
	cpuTdp           = 205 // in watts

)

func main() {
	wattPerThread := float64(cpuTdp) / cpuThreads                                                 // 3.79
	kiloWattHourPerThread := wattPerThread / 1000                                                 // 0.00379
	gramsOfCo2PerHourPerThread := kiloWattHourPerThread * gramsOfco2perkwh                        // 0.24
	gramsOfCo2PerThreadPerNS := gramsOfCo2PerHourPerThread / float64(time.Second/time.Nanosecond) // 0.00000000024296296296

	c := exec.Command("go", os.Args[1:]...)
	stdOut, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stdErr, err := c.StderrPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		io.Copy(os.Stderr, stdErr)
	}()

	r := regexp.MustCompile(`Benchmark(.*)-([0-8]+)(\s*)([0-9]+)(\s*)([0-9]+) ns/op(.*)`)

	go func() {
		b := bufio.NewReader(stdOut)
		for {
			row, err := b.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				panic(err)
			}

			if !r.MatchString(row) {
				fmt.Print(row)
				continue
			}

			x := r.FindStringSubmatch(row)

			nsPerOp, err := strconv.Atoi(x[6])
			if err != nil {
				panic(err)
			}

			gco2perop := float64(nsPerOp) * gramsOfCo2PerThreadPerNS

			fmt.Print(strings.TrimSpace(row))
			fmt.Printf("\t%.15f g CO2/op\n", gco2perop)
		}
	}()

	err = c.Run()
	if err != nil {
		panic(err)
	}
}
