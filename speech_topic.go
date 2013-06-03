package main

import "strings"

import "fmt"
import "os"

type ConfidenceResult struct {
	score float64
	response string
}

type SpeechTopic struct {
	name string
	hints []string
	response string
}

func (s *SpeechTopic) ScoreConfidence (input string) (ConfidenceResult) {
	var tokens = strings.Split(input, " ")

	var score = 0.0
	var scoreIncrement = float64(len(s.hints))

	for _, t := range tokens {
		for _, h := range s.hints {
			if strings.ToLower(t) == strings.ToLower(h) {
				score += scoreIncrement
			}
		}
	}

	fmt.Println(os.Stdout, s.name, ":", score * 100, "%")

	return ConfidenceResult { score, s.response }
}