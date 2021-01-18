package main
 
import (
    "fmt"
    // "time"
	"sync"
	"runtime"
)
 
// func cetak(name string) {
//     fmt.Println(name)
// }
 
// func cetak(name string, wg *sync.WaitGroup) {
//     wg.Done()
//     fmt.Println(name)
// }
func main() {
	
	// datanama := []string{"Gunawan", "Prasetyo",""}

    // datamhs := []string{"Didik", "Prabowo",
    //     "Charly", "Vah Houten", "Agus"}
 
    // // for _, v := range datamhs {
    //     time.Sleep(250 * time.Millisecond)
    //     go cetak(v)
	// }
	
    // for _, n := range datanama {
    //     time.Sleep(250 * time.Millisecond)
    //     go cetak(n)
	// }
	
	// var wg sync.WaitGroup
    // datamhs1 := []string{"Didik", "Prabowo",
    //     "Charly", "Vah Houten", "Agus"}
 
    // for _, v1 := range datamhs1 {
    //     wg.Add(1)
    //     go cetak(v1, &wg)
    // }
	// wg.Wait()
	runtime.GOMAXPROCS(10)
    var wg sync.WaitGroup
 
    wg.Add(10)
    go func() {
        defer wg.Done()
        fmt.Println("Ini fungsi Pertama ")
    }()
 
    go func() {
        defer wg.Done()
        fmt.Println("Ini fungsi Kedua ")
    }()
 
    wg.Wait()
}