package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result struct {
	result string
}

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

//162ms
func Google1_0(query string) (result []Result) {
	results := make([]Result, 0)
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

	return results
}

/**
Using go routines to parallelize
80ms
*/
func Google2_0(query string) (result []Result) {
	results := make([]Result, 0)
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return results
}

/**
Using timeout pattern to not wait more than 80ms to fetch all results
*/
func Google2_1(query string) (result []Result) {
	results := make([]Result, 0)
	c := make(chan Result)
	timeout := time.After(80 * time.Millisecond)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return results
		}
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google2_1("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
