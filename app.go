package main

import (
	"fmt"
	"io/ioutil"
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
	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func bored(w http.ResponseWriter, r *http.Request) {

	q := r.Body
	query,_ := ioutil.ReadAll(q)
	//querys, ok := r.URL.Query()["query"]
	//if !ok || len(querys[0]) < 1 {
	//	fmt.Fprintf(w, "Did not find that query!")
	//}
	//query := querys[0]
	fmt.Println("Number of participants queried: ", string(query))

	url := fmt.Sprintf("https://www.boredapi.com/api/activity?participants=%s", query)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Unable to make newrequest ",err)
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


	fmt.Fprintf(w, string(body))
}
