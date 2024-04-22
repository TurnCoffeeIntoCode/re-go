package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TurnCoffeeIntoCode/re-go/ssr"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		props := map[string]interface{}{
			"Name": "Rendered using ReGo",
		}
		queryString := r.URL.Query()
		if queryString.Get("name") != "" {
			props["Name"] = queryString.Get("name")
		}
		page := ssr.Page{
			Path:  r.URL.Path,
			Props: props,
		}
		err := page.Render(w)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(err.Error()))
		}
	})
	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
