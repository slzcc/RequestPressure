package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	c = flag.Int("c", 0, "Plz input client quantity")
	t = flag.Int("t", 0, "Plz input times quantity")
	u = flag.String("u", "", "Plz input url")
	o = flag.Bool("o", false, "output request status")
	oo = flag.Bool("oo", false, "output request body")
)

var (
	total    = 0.0
	about    = 0.0
	success  = 0.0
	failure  = 0.0
	use_time = 0.0
)

var wg sync.WaitGroup

func run(num int) {

	defer wg.Done()

	no := 0.0
	ok := 0.0

	for i := 0; i < num; i++ {

		sub_request_start_time := time.Now().UnixNano()

		resp, err := http.Get(*u)

		sub_request_end_time := time.Now().UnixNano()

		if err != nil {
			no += 1
			continue
		}

		defer resp.Body.Close()

		buf := bytes.NewBuffer(make([]byte, 0, 512))

		if *o || *oo != true {

			fmt.Printf("\nRequest StatusCode: %v %v, UseTime: %v %v, Request Success: %v, Request Failure: %v\n\n", resp.Status, resp.StatusCode, fmt.Sprintf("%.4f", float64(sub_request_end_time-sub_request_start_time)/1e9), "s", success, failure)

		}

		if *oo || *o != true {

			//headers := resp.Header

			//for k, v := range headers {
			//	fmt.Printf("k=%v, v=%v\n", k, v)
			//}

			//fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)
			//
			//fmt.Printf("resp Proto %s\n", resp.Proto)
			//
			//fmt.Printf("resp content length %d\n", resp.ContentLength)
			//
			//fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)
			//
			//fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)

			//fmt.Println(reflect.TypeOf(resp.Body))

			length, _ := buf.ReadFrom(resp.Body)

			fmt.Printf("\nRequest StatusCode: %v %v\nRequest Body Bytes Size: %v\nRequest Body Size: %v\nRequest UseTime: %v\nRequest Body Content: %v\n%v", resp.Status, resp.StatusCode, len(buf.Bytes()), length, fmt.Sprintf("%.4f", float64(sub_request_end_time-sub_request_start_time)/1e9), string(buf.Bytes()), "\n")
		}

		if resp.StatusCode != 200 {
			no += 1
			continue
		}

		ok += 1
		continue
	}

	success += ok
	failure += no
	total += float64(num)

}

func main() {

	start_time := time.Now().UnixNano()

	flag.Parse()

	if *c == 0 || *t == 0 || *u == "" {
		flag.PrintDefaults()
		return
	}

	fmt.Println("c:", *c, ",t:", *t)

	for i := 0; i < *c; i++ {
		wg.Add(1)
		go run(*t)
	}

	wg.Wait()
	end_time := time.Now().UnixNano()

	fmt.Println("PreTotal:", (*c)*(*t))
	fmt.Println("Total:", total)
	fmt.Println("Success:", success)
	fmt.Println("Failure:", failure)
	fmt.Println("SuccessRate:", fmt.Sprintf("%.2f", ((success/total)*100.0)), "%")
	fmt.Println("UseTime:", fmt.Sprintf("%.4f", float64(end_time-start_time)/1e9), "s")
}