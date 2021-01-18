package main
 
import (
    "fmt"
)
 
func main() {
 
    var a int = 10
    var b *int = &a
 
    fmt.Println("Alamat Memori variable b adalah", b)
    fmt.Println("Nilai dari variable b adalah", *b)
}