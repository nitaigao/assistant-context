package main

import "fmt"
import "os"
import "net/http"
import "net/url"

var conversation = Conversation { }

type ResponderHost struct {
	host string
	port string
}

var responders = [] ResponderHost {}

func respondToConversation(response string) {
	fmt.Println("Responding to conversation with", response)

	for _, v := range responders {
		var postUrl = "http://" + v.host + ":" + v.port + "/say"
		http.PostForm(postUrl, url.Values{"text":{response}})
		fmt.Println(postUrl)
	}	
}

func addToConversation(text string) {
	fmt.Println(os.Stdout, "Adding", text, "to the conversation")
	conversation.AddText(text)
	var confidence, response = conversation.Confidence()

	var minConfidence = 0.75
	if confidence > minConfidence {
		respondToConversation(response)
		conversation.StartNew()
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var text string = r.FormValue("text")
	addToConversation(text)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var responder = ResponderHost {r.FormValue("host"), r.FormValue("port")}
	responders = append(responders, responder)

	r.ParseForm()
	fmt.Println(os.Stdout, "Registered responder", r.Form["host"], "on port", r.Form["port"])
}

func main() {
	conversation.AddTopic( SpeechTopic { "Acknowledge", [] string { "Jarvis" }, "Yes Sir?" } )

	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}
