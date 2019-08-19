package main

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/jasonlvhit/gocron"
)

func main() {
    s := gocron.NewScheduler()
    s.Every(1).Minute().Do(task)
    <- s.Start()
}


func task() {

    url := "https://api.twilio.com/2010-04-01/Accounts/TWILIOSSID/Messages.json"

    payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"Body\"\r\n\r\n7\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"To\"\r\n\r\n+(TO PHONE NUMBER)\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"From\"\r\n\r\n+(FROM PHONE NUMBER)\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

    req, _ := http.NewRequest("POST", url, payload)

    req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
    req.Header.Add("account_sid", "XXX")
    req.Header.Add("Authorization", "Basic XXX")
    req.Header.Add("cache-control", "no-cache")
    req.Header.Add("Postman-Token", "XXX")

    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(res)
    fmt.Println(string(body))

    fmt.Println("Text was sent.")

}