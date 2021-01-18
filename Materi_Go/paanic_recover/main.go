// package main
// import (
//     "fmt"
// )
// // input
// func input(name string) {
//     fmt.Print("Masukkan Nama: ")
//     fmt.Scanln(&name)
 
//     if len(name) <= 3 {
//         panic("panjang teks harus lebih dari 3")
//     }
 
// }
// // cetak
// func cetak(nama *string) {
//     if nama == nil {
//         panic("Tidak boleh nil om")
//     }
// }
// func main() {
// 	var name string
//     input(name)
// 	fmt.Println("Belajar Golang ke 33\n.")
// 	cetak(&name)
// }
// ========================================= //
package main
 
import (
    "fmt"
)
 
func recoverCetak() {
    if r := recover(); r != nil {
        fmt.Println("error", r)
    }
}
 
func cetak(nama *string) {
    defer recoverCetak()
 	if nama == nil {
        panic(": Tidak boleh nil om")
    }
}
func main() {
    fmt.Println("Belajar Golang ke 33")
    cetak(nil)
}
 