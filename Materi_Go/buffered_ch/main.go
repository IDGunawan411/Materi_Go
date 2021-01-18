package main
 
import (
    "fmt"
)
 
func tulis(data []string, ch chan string) {
    for _, dt := range data {
        ch <- dt
    }
    close(ch)
}
 
func cetak(ch chan string) {
    for result := range ch {
        fmt.Println(result)
    }
}
 
func main() {
    data := []string{"PHP", "Golang", "Javascript", "Python"}
    ch := make(chan string, len(data))
    go tulis(data, ch)
    cetak(ch)
}