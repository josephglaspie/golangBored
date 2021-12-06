package main

import (
	"encoding/json"
	"fmt"
	"github.com/ecnepsnai/craigslist"
	"net/http"
	"os"
)

const (
	//query = "macbook"
	category = "sys"
	areaId = 308
	lat = 32.269112
	long = -95.307666
	dist = 100
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

http.HandleFunc("/computers",getItem)
http.HandleFunc("/",home)
http.ListenAndServe(":"+port,nil)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func getItem(w http.ResponseWriter, r *http.Request){

	querys, ok := r.URL.Query()["query"]
	if !ok || len(querys[0]) <1 {
		fmt.Fprintf(w,"Did not find that query!")
	}

	query := querys[0]

	fmt.Println("query: ",query)
	res, err := getCraigsListItem(category,query)
	if err !=nil{
		fmt.Printf("errored out ", err)
	}
	m := make(map[int]string)
	for _,v := range res {
		m[v.PostingID]= fmt.Sprintf("%v for $%d, located: %v",v.Title,v.Price, v.Location)
	}

	j, err := json.Marshal(m)
	if err != nil{
			panic(err)
	}
	fmt.Fprintf(w,string(j))
}

func getCraigsListItem(category, query string)([]craigslist.Result,error){
	results, err := craigslist.Search(category, query, craigslist.LocationParams{
		AreaID:         areaId,
		Latitude:       lat,
		Longitude:      long,
		SearchDistance: dist,
	})
	if err != nil {
		panic(err)
	}

	if len(results) == 0 {
		fmt.Printf("No results!")
	}

	return results,nil
}
