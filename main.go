package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){

	for i, arg := range os.Args[1:] {
		if(arg == "-I"){
			resp, _ := http.Get(os.Args[int(i+2)])
			for _, header := range resp.Header {
				fmt.Println(header)
			}
			break
		} else {
			resp, error := http.Get(arg)
			if error != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
				os.Exit(1)
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	}
}
