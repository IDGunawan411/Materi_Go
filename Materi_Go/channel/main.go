package main
 
import "fmt"

func cetak(ch chan int) {
    fmt.Println("ini dari gorotuine...")
    ch <- 10
}
// ------------------------
func penjumlahan(a, b int, ch chan int) {
    j := a + b
    ch <- j
}
func pengurangan(a, b int, ch chan int) {
    k := a - b
    ch <- k
} 
func main() {
    var a chan int
    if a == nil {
        a = make(chan int)
        fmt.Printf("Channel ini mempunya tipe :  %T\n", a)
	}
	
	fmt.Println("----------------------------")
	
    chdata := make(chan int)
    go cetak(chdata)
    nilai := <-chdata
    fmt.Println("nilai channel integer :", nilai)
	fmt.Println("ini dari fungsi main...")
	
	fmt.Println("----------------------------")
	
	chdata2 := make(chan int)
 
    data2 := []int{5, 6, 7, 3, 5, 6, 9}
 
    for _, y := range data2 {
        go func(y int) {
            chdata2 <- y
 
        }(y)
    }
 
    for i := 0; i < len(data2); i++ {
        fmt.Println(<-chdata2)
	}
	
	fmt.Println("----------------------------")

	j := make(chan int)
    k := make(chan int)
 
    go penjumlahan(10, 8, j)
    go pengurangan(10, 9, k)
 
    hasilJumlah, hasilKurang := <-j, <-k
 
    fmt.Printf("Hasil dari penjumlahan %d dengan %d adalah %d\n",
        10, 8, hasilJumlah)
 
    fmt.Printf("Hasil dari pengurangan %d dengan %d adalah %d\n",
		10, 8, hasilKurang)
		
}