	
package main
 
import (
    "fmt"
	"time"
)
 
func cetakHuruf1(ch chan string) {
    ch <- "Sedang Belajar Golang"
}
 
func cetakHuruf2(ch chan string) {
    ch <- "Golang itu Mudah"
 
}
// --------------
func lamarPT1(ch chan string) {
    ch <- "Berhasil Melamar di PT 1"
}
 
func lamarPT2(ch chan string) {
    ch <- "Berhasil Melamar di PT 2"
}

func main() {
    cHuruf1 := make(chan string)
    cHuruf2 := make(chan string)
 
    go func() {
        cetakHuruf1(cHuruf1)
    }()
    go func() {
        cetakHuruf2(cHuruf2)
    }()
 
    select {
    case c1 := <-cHuruf1:
        fmt.Println(c1)
    case c2 := <-cHuruf2:
        fmt.Println(c2)
	}
	fmt.Println("-----------------------------")

	lamar1 := make(chan string)
    lamar2 := make(chan string)
 
    go lamarPT1(lamar1)
    go lamarPT2(lamar2)
    time.Sleep(1 * time.Second)
    // select {
    // case l1 := <-lamar1:
    //     fmt.Println(l1)
    // case l2 := <-lamar2:
    //     fmt.Println(l2)
	// }

	for i := 0; i < 2; i++ {
		select {
		case l1 := <-lamar1:
			fmt.Println(l1)
		case l2 := <-lamar2:
			fmt.Println(l2)
		}
	}
}