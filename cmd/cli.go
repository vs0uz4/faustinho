package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/vs0uz4/faustinho/internal/loadtest"
	"github.com/vs0uz4/faustinho/internal/utils"
)

const errorMsg = "Error:"

func main() {
	url := flag.String("url", "", "URL to be tested")
	totalRequests := flag.Int("requests", 100, "Number of total requests")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent requests")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCLI Faustinho Stress Tester\n")
		fmt.Fprintf(os.Stderr, "---------------------------\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  --url		(string) URL to be tested. Required.\n")
		fmt.Fprintf(os.Stderr, "  --requests	(int) Number of total requests. Default: 100.\n")
		fmt.Fprintf(os.Stderr, "  --concurrency	(int) Number of concurrent requests. Default: 10.\n\n")
		fmt.Fprintf(os.Stderr, "Example:\n")
		fmt.Fprintf(os.Stderr, "  %s --url=http://example.com --requests=1000 --concurrency=10\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Parse()

	if err := utils.ValidateURL(*url); err != nil {
		fmt.Println(errorMsg, err)
		flag.Usage()
		os.Exit(1)
	}

	if err := utils.ValidatePositiveNumber(*totalRequests, "requests"); err != nil {
		fmt.Println(errorMsg, err)
		flag.Usage()
		os.Exit(1)
	}

	if err := utils.ValidatePositiveNumber(*concurrency, "concurrency"); err != nil {
		fmt.Println(errorMsg, err)
		flag.Usage()
		os.Exit(1)
	}
	startTime := time.Now()

	runner := loadtest.NewRunner(*url, *totalRequests, *concurrency)
	runner.Execute()

	endTime := time.Now()

	report := loadtest.GenerateReport(runner, startTime, endTime)
	report.Print()

	fmt.Println("Faustinho stress test completed!")
}
