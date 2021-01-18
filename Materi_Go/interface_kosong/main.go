package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
 
    var mp interface{}
 
    //tipe slice
    mp = []string{"Satu", "Dua"}
 
    var toSlice = strings.Join(mp.([]string), ", ")
    fmt.Println(toSlice)
	
 
    data := map[string]string{
        "Name":    "Didik Prabowo",
        "Address": "Wonogiri",
        "Website": "Kodingin.com",
    }
 
    mp = map[string]interface{}{
        "data": data,
    }
 
    fmt.Println(mp)
}