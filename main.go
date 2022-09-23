package main

import (
	"fmt"
	"os"
	"time"
)

type RunnerConfig struct {
	path  string
	count int
}

type BenchResult struct {
	InitTime  int64
	WriteTime int64
	ReadTime  int64
	FSize     int64
}

type TestRunner interface {
	Run(c RunnerConfig) (*BenchResult, error)
	Clean() error
}

func (r *BenchResult) String() string {
	total := r.InitTime + r.WriteTime + r.ReadTime
	return fmt.Sprintf("Total usage: %vms, Init usage: %vms, Read usage: %vms, Write usage: %vms, File size: %vM",
		total, r.InitTime, r.ReadTime, r.WriteTime, r.FSize/1024/1024)
}

func tempPath(prefix string) string {
	return os.TempDir() + prefix + "_" + randVal()
}

func runBench(runner TestRunner, name string, count int) error {
	tmpp := tempPath(name)
	// fmt.Println(tmpp)
	res, err := runner.Run(RunnerConfig{path: tmpp, count: count})
	defer runner.Clean()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %s\n", name, res)
	return nil
}

func main() {
	fmt.Printf("Now benchmark startat %s\n", time.Now())
	testCases := []int{10000, 100000, 1000000, 5000000}
	for _, n := range testCases {
		fmt.Printf("Now test %v count:\n", n)
		runBench(new(PudgeRunner), "pudge", n)
		runBench(new(BuntRunner), "buntdb", n)
		runBench(new(PogrebRunner), "pogreb", n)
	}
}
