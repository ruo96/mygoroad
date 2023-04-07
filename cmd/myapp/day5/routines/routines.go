package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.baidu.com/",
		"https://www.google.com/",
		"https://www.microsoft.com/",
		"https://www.apple.com/",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error downloading", url, ":", err)
				return
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading body from", url, ":", err)
				return
			}

			err = ioutil.WriteFile(url+".html", body, 0644)
			if err != nil {
				fmt.Println("Error writing file for", url, ":", err)
				return
			}

			fmt.Println("Downloaded and saved file for", url)
		}(url)
	}

	wg.Wait()
	fmt.Println("All downloads completed")
}
