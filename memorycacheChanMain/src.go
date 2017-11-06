package main

import (
	"../memoryCacheChan"
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
)

func getUrl(key string) (interface{}, error) {
	resp, err := http.Get(key)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	start := time.Now()
	cache := memoryCacheChan.New(getUrl)
	_, err := cache.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("error")
	} else {
		//fmt.Println(body)
	}
	end := time.Since(start)
	fmt.Println(end.Nanoseconds())

	start2 := time.Now()
	_, err2 := cache.Get("http://www.baidu.com")
	if err2 != nil {
		fmt.Println("error")
	} else {
		//fmt.Println(body2)
	}
	end2 := time.Since(start2)
	fmt.Println(end2.Nanoseconds())

	start3 := time.Now()
	_, err3 := cache.Get("http://www.baidu.com")
	if err3 != nil {
		fmt.Println("error")
	} else {
		//fmt.Println(body2)
	}
	end3 := time.Since(start3)
	fmt.Println(end3.Nanoseconds())
	
	cache.Close()
}
