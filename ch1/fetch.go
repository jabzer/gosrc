package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)


func main() {
	url := "http://www.baidu.com"
	resp,err :=http.Get(url)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"fet: %v\n",err)
		os.Exit(1)
	}
	b,err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr,"fet: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%s",b)
}