package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {
	// greeter("Hello") // first print Hello 5 times
	// go greeter("Hello") //  Go Routines // this not print world 5 times
	// greeter("World") // then print World 5 times

	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
	}

	wg.Add(len(websitelist))

	for _, web := range websitelist {
		getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i <= 5; i++ {
// 		time.Sleep(3 * time.Millisecond) // sleep for 3 millisecond to work routines properly
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPs in Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
