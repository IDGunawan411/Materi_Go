package fungsi

import "fmt"
import "net/http"
import "log"
func fungsi() {

	http.HandleFunc("/fungsi", 
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Fungsi World !!")
    })
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}

}