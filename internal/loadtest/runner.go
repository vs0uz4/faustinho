package loadtest

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

type Runner struct {
	URL             string
	TotalReqs       int
	Concurrency     int
	Results         []Result
	FailedRequests  int
	TimeoutRequests int
	httpClient      *http.Client
}

func NewRunner(url string, totalReqs, concurrency int) *Runner {
	return &Runner{
		URL:         url,
		TotalReqs:   totalReqs,
		Concurrency: concurrency,
		Results:     make([]Result, 0, totalReqs),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (r *Runner) Execute() {
	var wg sync.WaitGroup
	requests := make(chan struct{}, r.Concurrency)
	resultLock := sync.Mutex{}

	for i := 0; i < r.TotalReqs; i++ {
		requests <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			resp, err := r.httpClient.Get(r.URL)
			result := Result{Duration: time.Since(start)}

			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					result.StatusCode = -2
					resultLock.Lock()
					r.TimeoutRequests++
				} else {
					result.StatusCode = -1
					resultLock.Lock()
					r.FailedRequests++
				}
				r.Results = append(r.Results, result)
				resultLock.Unlock()
				<-requests
				return
			}

			result.StatusCode = resp.StatusCode
			resp.Body.Close()
			resultLock.Lock()
			r.Results = append(r.Results, result)
			resultLock.Unlock()
			<-requests
		}()
	}

	wg.Wait()
}
