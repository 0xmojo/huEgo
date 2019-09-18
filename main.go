package main

import (
    //"encoding/json"
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "os"
)
func main() {
	var ip string = os.Args[1]
	var port string = "80"

	var url string = "http://" + ip + ":" + port;
	fmt.Println(url)
	resp, err := http.Get(url + "/api")
	if err != nil {
		log.Fatal(err);
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	//decode json and so on
	defer resp.Body.Close()
}
