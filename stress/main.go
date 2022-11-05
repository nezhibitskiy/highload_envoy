package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//wg := sync.WaitGroup{}

	for i := 0; i < 1000000; i++ {
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		resp, err := http.Get("http://localhost:5000/api/user/user1/profile")
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status, resp.Header.Get("hostname"))
		resp.Body.Close()
		//}()
	}

	//wg.Wait()
}
