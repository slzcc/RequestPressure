package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"net/url"
)

var (
	c = flag.Int("c", 0, "Plz input client quantity")
	t = flag.Int("t", 0, "Plz input times quantity")
	u = flag.String("u", "", "Plz input url")
	o = flag.Bool("o", false, "output request status")
	oo = flag.Bool("oo", false, "output request body")
)

var (
	urls      = string("https://kaipan.tahoecn.com/opening/app/house/getCanOrder")
	total    = 0.0
	about    = 0.0
	success  = 0.0
	failure  = 0.0
	use_time = 0.0
)

var wg sync.WaitGroup


func run(num int, data, response string) {
	defer wg.Done()

	no := 0.0
	ok := 0.0

	for i := 0; i < num; i++ {

		sub_request_start_time := time.Now().UnixNano()

		urlStr := urls + "?AES_DATA=" + url.QueryEscape(data)

		resp, err := http.Post(urlStr, "application/json", nil)

		sub_request_end_time := time.Now().UnixNano()

		if err != nil {
			no += 1
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			body = nil
			continue
		}

		if *o {

			fmt.Printf("Request StatusCode: %v %v, UseTime: %v %v, Request Success: %v, Request Failure: %v\n\n", resp.Status, resp.StatusCode, fmt.Sprintf("%.4f", float64(sub_request_end_time-sub_request_start_time)/1e9), "s", success, failure)

		}else if *oo {

			fmt.Printf("Response Body: %v\n,Request StatusCode: %v %v, UseTime: %v %v, Request Success: %v, Request Failure: %v, Request Body Content: %v\n\n", response, resp.Status, resp.StatusCode, fmt.Sprintf("%.4f", float64(sub_request_end_time-sub_request_start_time)/1e9), "s", success, failure, string(body))

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

		resp, err := http.Get(*u)

		if err != nil {

			continue
		}

		defer resp.Body.Close()

		buf := bytes.NewBuffer(make([]byte, 0, 512))
		length, _ := buf.ReadFrom(resp.Body)

		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))

		if length == 0 {
			fmt.Println(length)
		}
		var f interface{}
		json.Unmarshal(buf.Bytes(), &f)

		for key, value := range f.(map[string]interface{}) {
			switch key {
				case "data":
					jsonStr, err := json.Marshal(value)
					if err != nil {
						fmt.Println("error:", err)
					}
					go run(*t, string(jsonStr), string(buf.Bytes()))
			}
		}

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