package main
import "fmt"
import "net/http"
import "log"
func main() {

	http.HandleFunc("/", 
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Hello World")
    })
    http.HandleFunc("/hello", 
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "World Hello 2")
	})
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}

}