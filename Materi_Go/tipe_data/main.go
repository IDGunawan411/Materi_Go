package main
 
import (
    "fmt"
)
func main() {
    var nama, alamat, sekolah string
    var umur int
 
    nama = "Udin Raharjo"
    alamat = "Solo"
    sekolah = "SMA N 1 Solo"
    umur = 21
    tinggibadan := 156
    beratbadan := 50
    fmt.Println("=== Biodata ===")
    fmt.Println("Nama : " + nama)
    fmt.Println("Alamat : " + alamat)
    fmt.Println("Sekolah : " + sekolah)
    fmt.Println("Umur : ", umur)
    fmt.Println("Tinggi Badan", tinggibadan)
    fmt.Println("Berat Badan", beratbadan)
}