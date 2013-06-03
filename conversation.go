package main

type Conversation struct {
	input string
	topics [] SpeechTopic
}

func (c *Conversation) AddText(text string) { 
	if len(c.input) > 0 {
		c.input += " "
	}
	c.input += text
}

func (c* Conversation) AddTopic(topic SpeechTopic) {
	c.topics = append(c.topics, topic)
}

func (c *Conversation) Confidence() (float64, string) {
	var confidenceScores = [] ConfidenceResult { }
	for _, v  := range c.topics {
		var confidenceScore = v.ScoreConfidence(c.input)
		confidenceScores = append(confidenceScores, confidenceScore)
	}

	if len(confidenceScores) > 0 {
		var confidenceScore = confidenceScores[len(confidenceScores) - 1]
		return confidenceScore.score, confidenceScore.response
	}

	return 0, ""
}

func (c *Conversation) StartNew() {
	c.input = ""
}