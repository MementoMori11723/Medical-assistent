package gpt

import (
	"Medical-assistent/config"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type image struct {
	Url string `json:"url"`
}

type content struct {
	Type      string `json:"type"`
	Text      string `json:"text,omitempty"`
	Image_url image  `json:"image_url,omitempty"`
}

type message struct {
	Role    string    `json:"role"`
	Content []content `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func New(data string, img string) string {
	url, apiKey, Type := config.GetGptData()
	var request Request
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	request = Request{
		Model: "deepseek-chat",
		Messages: []message{
			{
				Role: "system",
				Content: []content{
					{
						Type: "text",
						Text: Type,
					},
				},
			},
			{
				Role: "user",
				Content: []content{
					{
						Type: "text",
						Text: data,
					},
				},
			},
		},
	}

	if img != "" {
		request = Request{
			Model: "deepseek-chat",
			Messages: []message{
				{
					Role: "system",
					Content: []content{
						{
							Type: "text",
							Text: Type,
						},
					},
				},
				{
					Role: "user",
					Content: []content{
						{
							Type: "text",
							Text: data,
						},
						{
							Type: "image_url",
							Image_url: image{
								Url: img,
							},
						},
					},
				},
			},
		}

	}

	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshaling request: %v", err)
	}

	req, err := http.NewRequest(
		"POST", url,
		bytes.NewBuffer(jsonReq),
	)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		log.Fatalf("Request failed with status %d: %s", res.StatusCode, string(resBody))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var response Response
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	return response.Choices[0].Message.Content
}
