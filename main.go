package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

var (
    // \d is a number
    reQQEmail = `(\d+)@qq.com`
)

// Crawl mailbox
func GetEmail() {
    // 1. Go to the website to get the data
    url := "https://toidicodedao.com/2015/02/05/callback-trong-javascript/"
    resp, err := http.Get(url)
    if err != nil{
        fmt.Println(err.Error())
        return
    }
    defer resp.Body.Close()
    // 2. Read page content
    pageBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		fmt.Println(err.Error())
		return
	}
    pageStr := string(pageBytes)

    // fmt.Println(pageStr)
	// // 3. Filter data, filter qq mailbox
    re := regexp.MustCompile(reQQEmail)

    results := re.FindAllStringSubmatch(pageStr, -1)
    fmt.Println(results)

}

func main() {
    GetEmail()
}
