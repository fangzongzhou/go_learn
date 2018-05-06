package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getRespose(url string) string {
	if !strings.Contains(url,"http://"){
		url="http://"+url
		fmt.Print(url)
	}
	response, _ := http.Get(url)
	b,_:=ioutil.ReadAll(response.Body)


	return string(b)
}

func main() {
	for _, url := range os.Args[1:] {

		fmt.Printf("%s", getRespose(url))
	}
}
