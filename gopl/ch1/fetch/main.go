package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// ex1.8 version
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v", err)
			os.Exit(1)
		}

		// origin version
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n Code: %s", b, resp.Status)

		/*
			// ex1.7 version
			_, err = io.Copy(os.Stdout, resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %v", err)
				os.Exit(1)
			}
		*/
	}
}
