package main

import (
	"bytes"
	"encoding/json"
	// "github.com/Youknow2509/discord_webhook/internal/model"
	"log"
	"net/http"
	"time"
)

func main() {
	urlWebhook := "https://discord.com/api/webhooks/1363227925687701706/IWxsdQkrWpVLt9qEUflvksk2PYVmfnrGM8FmfJsnc7RfJ-izZxyUTTtlZKWBIVvwcpzb"

	// message := model.Message{
	// 	Username: "Test user",
	// 	Content:  "Test content",
	// }

	in := []byte(`{
"content": "Hello, World!",
  "embeds": [{
    "title": "Hello, Embed!",
    "description": "This is an embedded message."
  }]
}
`)
	var obj interface{}
	err := json.Unmarshal(in, &obj)
	if err != nil {
		panic(err)
	}

	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(obj)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(payload)

	req, err := http.NewRequest(
		http.MethodPost,
		urlWebhook,
		payload,
	)
	if err != nil {
		log.Fatal("Error whening post ")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("client: error making http request: %s\n", err)
	}
	log.Println(res)
}
