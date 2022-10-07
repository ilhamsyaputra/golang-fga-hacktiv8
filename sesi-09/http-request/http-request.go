package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v \n", res.Body)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)
}

func postRequest() {
	data := map[string]interface{}{
		"title":  "M. Ilham Syaputra",
		"body":   "Future Engineer",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data) //mengubah data ke JSON
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func main() {
	// getRequest()
	postRequest()
}
