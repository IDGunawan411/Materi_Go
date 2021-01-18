	
package main
 
import (
    "fmt"
)
 
func main() {
	
	fmt.Println("---------------- Operator ----------------")

    //Set Nilai Awal
    angka1 := 10
    angka2 := 5
 
    //Operator Penjumlahan
    hasiljumlah := angka1 + angka2
    fmt.Println(fmt.Sprintf("Hasil penjumlahan dari %d dan %d adalah %d Bos..", angka1, angka2, hasiljumlah))
 
    //Operator Pengurangan
    hasilkurang := angka1 - angka2
    fmt.Println(fmt.Sprintf("Hasil Pengurangan dari %d dan %d adalah %d Bos..", angka1, angka2, hasilkurang))
 
    //Operator Perkalian
    hasilkali := angka1 * angka2
    fmt.Println(fmt.Sprintf("Hasil Perkalian dari %d dan %d adalah %d Bos..", angka1, angka2, hasilkali))
 
    //Operator Pembagian
    hasilbagi := angka1 / angka2
    fmt.Println(fmt.Sprintf("Hasil Pembagian dari %d dan %d adalah %d Bos..", angka1, angka2, hasilbagi))
 
    //Operator Perkalian
    hasilsisa := angka1 % angka2
    fmt.Println(fmt.Sprintf("Hasil Sisa Hasil Bagi dari %d dan %d adalah %d Bos..", angka1, angka2, hasilsisa))
	
	fmt.Println("---------------- Perbandingan ----------------")

	angka3 := angka1 > angka2
	fmt.Println(angka3)

	fmt.Println("---------------- Logika ----------------")

    a := true
    b := false
 
    c := a && b
    fmt.Println(c)

	fmt.Println("---------------- Penugasan ----------------")

	pen := 5
	pen += 4

	fmt.Println(pen)

	fmt.Println("---------------- Bitwise ----------------")

	bit_a := 00000101
    bit_b := 00000100
 
    bit_c := bit_a & bit_b
 
    fmt.Println(bit_c)
}