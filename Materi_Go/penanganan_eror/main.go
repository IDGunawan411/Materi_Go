package main
 
import (
    "fmt"
    "os"
)
 
func main() {
	// file, err := os.Open("note.txt")
	file, err := os.Open("noted.txt")
 
    if err != nil {
        fmt.Println("Terjadi Kesalahan om yaitu", err.Error())
    }
 
    fmt.Println(file)
 
}