package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"sort"
	"syscall"
	"time"
)

var list = sort.IntSlice([]int{1, 2, 3})

func Goroutine() {
	go list.Sort()
}

func Announce(message string, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		println(message)
	}()
}

func Channel() {
	ch := make(chan int)
	go func() {
		list.Sort()
		ch <- 1
	}()
	// do something
	<-ch
}

var sem = make(chan int, 10)

func process(r *http.Request) {
	// long task
}

func handle(r *http.Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}

func Serve1(queue chan *http.Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

func Serve2(queue chan *http.Request) {
	for req := range queue {
		req := req // rebind for use in goroutine
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func handle2(queue chan *http.Request) {
	for req := range queue {
		process(req)
	}
}

func Serve3(clientRequests chan *http.Request, quit chan bool) {
	for i := 0; i < 10; i++ {
		go handle2(clientRequests)
	}
	<-quit
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

func request() {
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	//clientRequests <- request
	fmt.Printf("answer: %d\n\n", <-request.resultChan)
}

func handle3(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

type Vector []float64

// DoSome Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

func (v Vector) Op(f float64) float64 {
	return f
}

var numCPU = runtime.GOMAXPROCS(0)

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // Wait for one task to complete.
	}
	// All done.
}

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

func client() {
	var b *bytes.Buffer
	select {
	case b = <-freeList:
		// Got one; nothing more to do.
	default:
		// None free, so allocate a new one.
		b = new(bytes.Buffer)
	}
	//load(b)
	serverChan <- b
}

func server() {
	for {
		b := <-serverChan
		//process(b)
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default:
			// Free list full, just carry on.
		}
	}
}

func err() {
	for try := 0; try < 2; try++ {
		_, err := os.Create("file")
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			//deleteTempFiles()
			continue
		}
		return
	}
}

// CubeRoot A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
	z := x / 3 // Arbitrary initial value
	for i := 0; i < 1e6; i++ {
		prevZ := z
		z -= (z*z*z - x) / (3 * z * z)
		if veryClose(z, prevZ) {
			return z
		}
	}
	panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

func veryClose(a float64, b float64) bool {
	return math.Abs(a-b) < 1e-10
}
