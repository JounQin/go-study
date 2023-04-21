package main

import (
	"log"
	"os"
	"testing"
)

func TestJob_Print(t *testing.T) {
	var job = NewJob("command %s", log.New(os.Stderr, "Job: ", log.Ldate))
	job.Print("test")
	job = &Job{"%s command", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("test")
}
