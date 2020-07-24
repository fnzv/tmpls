package main

import (
	"fmt"
	"os"
	"log"
	"bytes"
	"strings"
    "net/http"
)

var ipreal string
func telegramNotify(msg string) {


	if os.Getenv("TGBOT_TOKEN") != "" && os.Getenv("TGBOT_CHATID") != "" {
    url := "https://api.telegram.org/"+os.Getenv("TGBOT_TOKEN")+"/sendMessage"
	var jsonStr = []byte(`{"chat_id": `+os.Getenv("TGBOT_CHATID")+`, "text": "`+msg+`"}`)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
} else {

	fmt.Println("Telegram env vars are not set... skipping notification")
}
}

func headers(w http.ResponseWriter, req *http.Request) {
	f, err := os.OpenFile("tracker.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
			log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	for k, v := range req.Header {
		log.Println("Header detected ", k, v)
		//telegramNotify("Header detected "+ k+ ": "+ v)
		if k == "X-Original-Forwarded-For" {
			ipreal = fmt.Sprintf(strings.Join(v," "))
		}
    }

	ua := req.Header.Get("User-Agent")
	// if only one expected 
	uri := req.URL.Query().Get("uri")
	if uri != "" {
		log.Println("IP "+ipreal+" just opened "+uri)
		log.Println("User-agent: "+ua)
		telegramNotify("IP "+ipreal+" just opened "+uri)
		telegramNotify("User-agent: "+ua)
		//redirect to original host
		http.Redirect(w, req, "http://"+uri, 301)
	}
}

func main() {

    http.HandleFunc("/", headers)

    http.ListenAndServe(":80", nil)
}