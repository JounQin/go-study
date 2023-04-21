package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Stringer interface {
	String() string
}

func Type(value interface{}) string {
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return fmt.Sprintf("%v", value)
}

type Counter int

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	_, _ = fmt.Fprintf(w, "counter = %d ", *c)
}

type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch <- r
	_, _ = fmt.Fprint(w, "notification sent")
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, os.Args)
}

func httpHandle() {
	http.Handle("/notify", http.HandlerFunc(ArgServer))
}

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Print(args ...interface{}) {
	job.Logger.Printf(job.Command, args...)
}
