package main
 
import (
    "fmt"
)
 
func main() {
    nilai := 10
 
    if nilai > 7 {
        fmt.Println("Anda Lulus dengan Nilai", nilai)
    } else {
        fmt.Println("Anda Tidak Lulus dengan Nilai", nilai)
	}
	fmt.Println("--------------------------------------")
	var s = "Indonesia"
    x := false
    if x {
        fmt.Println("Anda berada di", s)
    } else {
        fmt.Println("Anda tidak berada di", s)
	}
	fmt.Println("-------------------- Switch Case ------------------")
	
	nilai_a := 70
 
    switch nilai_a {
    case 50:
        fmt.Println("Nilai Anda C")
 
    case 70:
        fmt.Println("Nilai Anda B")
    case 100:
        fmt.Println("Nilai Anda A")
    default:
        fmt.Println("Tidak ada kategori")
	}
	fmt.Println("----------- Switch Case Falltrough --------")
	
    switch nilai_a {
    case 50:
        fmt.Println("Nilai Anda C")
 
    case 70:
        fmt.Println("Nilai Anda B")
        fallthrough
    case 100:
        fmt.Println("Nilai Anda A")
 
    default:
        fmt.Println("Tidak ada kategori")
	}
	
	
}
