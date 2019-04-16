package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

func request(kind string) string {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	return fmt.Sprintf("%s, ", kind)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	c := make(chan string)
	go func() { c <- request("foo") }()
	go func() { c <- request("bar") }()
	go func() { c <- request("baz") }()
	go func() { c <- request("qux") }()

	all := make(chan string)
	go func() {
		var res string
		for i := 0; i < 4; i++ {
			select {
			case r := <-c:
				res += r
				fmt.Printf("all: %s\n", res)
			case <-ctx.Done():
				return
			}
		}
		all <- res
	}()

	select {
	case response := <-all:
		fmt.Fprint(w, response)
	case <-ctx.Done():
		fmt.Fprint(w, "timeout")
	}
	return
}

func main() {
	server := httptest.NewServer(http.HandlerFunc(handleRoot))
	defer server.Close()

	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	res, _ := http.Get(server.URL)
	elapsed := time.Since(start)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Response: %s (%s)\n", body, elapsed)
	time.Sleep(1200 * time.Millisecond)
}
