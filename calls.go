package main

import (
	"fmt"
	"gopkg.in/go-resty/resty.v2"
	"time"
	"sync"
	"strconv"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()
	
	wg := sync.WaitGroup{}

	client := resty.New()
	
	Range := []int{1, 2, 3, 4, 5}
	for _, id := range Range {

		wg.Add(1)
		go func(id int) {

      // https://www.blitter.se/utils/basic-authentication-header-generator/ for Authorzation header
			resp, err := client.
						R().
						EnableTrace().
						SetHeader("Authorization", "Basic YWRtaW46YWRtaW4=").
						SetHeader("Content-Type", "application/yang-data+json").
						Post("http://localhost:8080/restconf/operations/tailf-ncs:devices/tailf-ncs:device=R" + strconv.Itoa(id) + "/tailf-ncs:sync-from")
			
			fmt.Println("Response Info:")
			fmt.Println("  Error      :", err)
			fmt.Println("  Status Code:", resp.StatusCode())
			fmt.Println("  Status     :", resp.Status())
			fmt.Println("  Proto      :", resp.Proto())
			fmt.Println("  Time       :", resp.Time())
			fmt.Println("  Received At:", resp.ReceivedAt())
			fmt.Println("  Body       :\n", resp)
			fmt.Println()
			
			wg.Done()
		}(id)

	}
		
	wg.Wait()
}
