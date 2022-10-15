package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	u1 := "/"
	u2 := "/bwp"
	var urls = []string{u1, u2}
	log.Printf("start http server: %+v\n", urls)

	http.HandleFunc(u1, handler)
	http.HandleFunc(u2, printBody)
	err := http.ListenAndServe(":8895", nil)
	if err != nil {
		fmt.Println("start http server error: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	fmt.Fprintf(w, "hhhh hello\n")
}

func printBody(w http.ResponseWriter, r *http.Request) {
	bs, _ := ioutil.ReadAll(r.Body)
	body := string(bs[:100])
	fmt.Println(body)

	// w.Header().Add("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Methods", "*") //允许请求方法
	// w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	fmt.Fprintf(w, "print bwp\n")
}
