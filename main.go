package main

import "fmt"
import "strconv"
import "net/http"
import "os"
import "time"

func main(){

	url := os.Args[1]
	status,_ := strconv.Atoi(os.Args[2])
	threads,_ := strconv.Atoi(os.Args[3])

	cnt := 0
	lastcnt := 0
	
	for i := 0; i < threads; i++ {
		go func(){
			for {
				resp, err := http.Get(url)
				if err == nil {
					if resp.StatusCode == status {
						cnt += i
					} else {
						fmt.Printf("unexpected status: %d\n", resp.Status)
					}
					resp.Body.Close()
				} else {
					fmt.Printf("request error: %s\n", err)
				}
			}
		}()
	}
	
	for {
		time.Sleep(time.Second)
		diff := cnt-lastcnt
		fmt.Printf("resp/sec: %d\n", diff)
		lastcnt = cnt
	}

}
