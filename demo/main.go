package main

import (
	"fmt"
	"net/http"

	"github.com/arthurkiller/perfm"
)

type job struct {
	// job private data
	url string
}

func (j *job) String() string {
	return fmt.Sprintf("demo job: send GET to %s", j.url)
}

// Copy will called in parallel
func (j *job) Copy() (perfm.Job, error) {
	jc := *j
	return &jc, nil
}

func (j *job) Pre() error {
	// do pre job
	return nil
}
func (j *job) Do() error {
	// do benchmark job
	_, err := http.Get(j.url)
	return err
}
func (j *job) After() {
	// do clean job
}

func main() {
	// start it directly!
	j := &job{url: "http://www.baidu.com"}
	perfm.Regist(j)
	perfm.Start()

	// or control it yourself!
	p := perfm.New(perfm.WithBinsNumber(15), perfm.WithParallel(5), perfm.WithDuration(5))
	p.Regist(j)
	p.Start()
}
