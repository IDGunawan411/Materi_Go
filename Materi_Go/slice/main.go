package main
 
import (
    "fmt"
)
 
func main() {
    var mahasiswa []string
    mahasiswa = []string{"Didik Prabowo", "Charly Van Houten"}
 
	fmt.Println("Data Mahasiswa => ", mahasiswa)
	
    fmt.Println("Data Mahasiswa ke 0 sampai 1 => ", mahasiswa[0:1])
	fmt.Println("Data Mahasiswa ke 0 sampai 2 => ", mahasiswa[0:2])
	
    fmt.Println("Jumlah Mahasiswanya adalah => ", len((mahasiswa)))

    fmt.Println("--------- Salin nilai slice -----------------")

	biologi := []string{"Didik Prabowo", "Charly Van Houten"}
	fisika := make([]string, len(biologi))
	
    copy(fisika, biologi)
    fmt.Println(fisika)
	
	fmt.Println("--------- Element slice -----------------")
	
    mahasiswalama := []string{"Didik Prabowo", "Charly Van Houten"}
    mahasiswabaru := "Udin Kumalasari"
    mahasiswa_a := append(mahasiswalama, mahasiswabaru)
    fmt.Println(mahasiswa_a)
 
}