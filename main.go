package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"servergo/pkg"
)

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content       Content `json:"content"`
	FinishReason  string  `json:"finishReason"`
	Index         int     `json:"index"`
	SafetyRatings []any   `json:"safetyRatings"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text string `json:"text"`
}

func main() {
	fmt.Println("main")

	geminiKey := pkg.Getval_five("geminiKey")
	if geminiKey == "YOUR_API_KEY" {
		log.Fatal("Please replace YOUR_API_KEY with your actual Gemini API key.")
	}

	endPoint := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", geminiKey)

	// This is the request body you send TO the API
	requestPayload := map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{
						"text": "whos the current president of the USA",
					},
				},
			},
		},
	}
	requestBodyBytes, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatal("Error marshalling request body:", err)
	}

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("API request failed with status: %s (Code: %d)", resp.Status, resp.StatusCode))
	}

	var apiResponse GeminiResponse

	derr := json.NewDecoder(resp.Body).Decode(&apiResponse)
	if derr != nil {

		panic(derr)
	}

	//parsing AI response
	if len(apiResponse.Candidates) > 0 &&
		len(apiResponse.Candidates[0].Content.Parts) > 0 {
		fmt.Println("Gemini's reply:", apiResponse.Candidates[0].Content.Parts[0].Text)
	} else {
		fmt.Println("Could not find text in Gemini's response. Full response:")
		responseBytes, _ := json.MarshalIndent(apiResponse, "", "  ")
		fmt.Println(string(responseBytes))
	}
}
