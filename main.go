package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	
	if !strings.Contains(strings.ToLower(body.Message.Text), "testing") {
		return
	}

	
	if err := test(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}


type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func test(chatID int64) error {
	
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "@GPCMR",
	}
	
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	
	res, err := http.Post("https://api.telegram.org/bot1136074705:AAFrV0GgCNZldeVy-hTcgV6pfnvkb3yd6yg/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// FInally, the main funtion starts our server on port 3000
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
