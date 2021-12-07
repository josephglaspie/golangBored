package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
	"os"
)

var decoder  = schema.NewDecoder()


type twilioResponse struct {
//	ToCountry     string `schema:"ToCountry"`
//	ToState       string `schema:"ToState"`
//	SmsMessageSid string `schema:"SmsMessageSid"`
//	NumMedia      string `schema:"NumMedia"`
//	ToCity        string `schema:"ToCity"`
//	FromZip       string `schema:"FromZip"`
//	SmsSid        string `schema:"SmsSid"`
//	FromState     string `schema:"FromState"`
//	SmsStatus     string `schema:"SmsStatus"`
//	FromCity      string `schema:"FromCity"`
	Body          string `in:"form=Body"`
//	FromCountry   string `schema:"FromCountry"`
//	To            string `schema:"To"`
//	ToZip         string `schema:"ToZip"`
//	NumSegments   string `schema:"NumSegments"`
//	MessageSid    string `schema:"MessageSid"`
//	AccountSid    string `schema:"AccountSid"`
//	From          string `schema:"From"`
//	ApiVersion    string `schema:"ApiVersion"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/bored", bored)
	http.HandleFunc("/", home)
	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	begining := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Message>`
	middle := "Hello World!"
	end := `</Message>
</Response>
`
	out := fmt.Sprint(begining + middle + end)
	fmt.Fprintf(w, out)
}

func bored(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("Body")

	fmt.Println("Number of participants queried: ", query)

	url := fmt.Sprintf("https://www.boredapi.com/api/activity?participants=%s", query)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Unable to make newrequest ", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to client.do ", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("unable to readall ", err)
		return
	}

	twimlgo := twiml(string(body))

	fmt.Fprintf(w, twimlgo)
}

func twiml(input string) string {
	begining := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Message>`
	middle := input
	end := `</Message>
</Response>
`
	out := fmt.Sprint(begining + middle + end)
	return out
}
