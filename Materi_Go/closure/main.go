package main
 
import (
    "fmt"
)
 
func todo() func() int {
    i := 2
    return func() int {
        i++
        return i
    }
}

func main() {
 
    go1 := func() {
        fmt.Println("Belajar Golang!")
    }
    go2 := func() {
        fmt.Println("Belajar PHP!")
    }
    go1()
	go2()
	
    dataTodo := todo()
    fmt.Println("Nilai toDo adalah", dataTodo())
}