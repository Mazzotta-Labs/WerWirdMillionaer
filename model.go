package main

type Questions struct {
	Questions []Question
}

type Question struct {
	Id         string   `json:"id"`
	Question   string   `json:"question"`
	Difficulty int      `json:"difficulty"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}
