package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/bored", bored)
	http.HandleFunc("/", home)
	err := http.ListenAndServe("127.0.0.1:"+port, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		log.Fatal("Server failed to load ",err)
	}
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
