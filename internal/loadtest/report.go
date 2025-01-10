package loadtest

import (
	"fmt"
	"time"
)

type Report struct {
	TotalRequests      int
	SuccessfulRequests int
	FailedRequests     int
	TimeoutRequests    int
	StatusDistribution map[int]int
	TotalTime          time.Duration
	MinTime            time.Duration
	MaxTime            time.Duration
	AvgTime            time.Duration
}

func GenerateReport(runner *Runner, startTime time.Time, endTime time.Time) Report {
	statusDist := make(map[int]int)
	var totalTime time.Duration
	minTime := time.Duration(1<<63 - 1)
	maxTime := time.Duration(0)
	successfulRequests := 0

	for _, result := range runner.Results {
		statusDist[result.StatusCode]++
		totalTime += result.Duration

		if result.Duration < minTime {
			minTime = result.Duration
		}
		if result.Duration > maxTime {
			maxTime = result.Duration
		}

		if result.StatusCode == 200 {
			successfulRequests++
		}
	}

	statusDist[-2] = runner.TimeoutRequests
	statusDist[-1] = runner.FailedRequests

	totalRequests := runner.RequestsSent
	avgTime := time.Duration(0)
	if totalRequests > 0 {
		avgTime = totalTime / time.Duration(totalRequests)
	}

	return Report{
		TotalRequests:      totalRequests,
		SuccessfulRequests: successfulRequests,
		FailedRequests:     runner.FailedRequests,
		TimeoutRequests:    runner.TimeoutRequests,
		StatusDistribution: statusDist,
		TotalTime:          endTime.Sub(startTime),
		MinTime:            minTime,
		MaxTime:            maxTime,
		AvgTime:            avgTime,
	}
}

func (r Report) Print() {
	fmt.Println("\nFaustinho Stress Test Report")
	fmt.Println("------------------------------")
	fmt.Printf("Total Requests: %d\n", r.TotalRequests)
	fmt.Printf("Successful Requests (200): %d\n", r.SuccessfulRequests)
	fmt.Println("Status Code Distribution:")
	for code, count := range r.StatusDistribution {
		if code == -1 {
			fmt.Printf("  Client Errors: %d\n", count)
		} else if code == -2 {
			fmt.Printf("  Timeouts: %d\n", count)
		} else {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
	fmt.Printf("Total Time: %s\n", r.TotalTime)
	fmt.Printf("Min Time per Request: %s\n", r.MinTime)
	fmt.Printf("Max Time per Request: %s\n", r.MaxTime)
	fmt.Printf("Avg Time per Request: %s\n", r.AvgTime)
}
