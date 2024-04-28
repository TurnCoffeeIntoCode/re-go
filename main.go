package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TurnCoffeeIntoCode/re-go/ssr"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numbers := make([]string, 0, 100)
		for i := 0; i <= 100; i++ {
			numbers = append(numbers, "Number "+strconv.Itoa(i))
		}
		props := make(map[string]interface{})
		props["numbers"] = numbers
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
