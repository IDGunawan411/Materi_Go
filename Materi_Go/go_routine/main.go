package main
import (
    "fmt"
    "time"
)
 
func say() {
    fmt.Println("Halo ini berasal dari goroutine baru")
}
// --------------

func cetakAngka() {
    angka := []int{1, 2, 3, 4, 5}
    for x := range angka {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(angka[x])
    }
}
 
func cetakHuruf() {
    huruf := []string{"a", "b", "c", "d", "e"}
    for x := range huruf {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(huruf[x])
    }
}

func main() {
    go say()
    time.Sleep(1 * time.Second)
	fmt.Println("Ini Fungsi Utama")

	fmt.Println("-------------------")
	
	go cetakAngka()
    go cetakHuruf()
    time.Sleep(3000 * time.Millisecond)
}