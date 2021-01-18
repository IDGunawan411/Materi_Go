package main
 
import (
    "fmt"
)
 
// Profil
type Profil struct {
    ID      int
    Name    string
    Age     int
    Address string
}
 
type (
    Education struct {
        SMA string
        SMP string
    }
 
    Profil2 struct {
        ID        int
        Name      string
        Age       int
        Address   string
        Education Education
    }
)
func main() {
    var profil Profil
 
    profil.ID = 1
    profil.Name = "Didik Prabowo"
    profil.Age = 20
    profil.Address = "Wuryantoto, Wonogiri"
 
    fmt.Println(profil)
    fmt.Println("===========================")
    fmt.Printf("Nama \t: %v\n", profil.Name)
    fmt.Printf("Umur \t: %d\n", profil.Age)
    fmt.Printf("Alamat \t: %v\n", profil.Address)
    fmt.Println("===========================")

	profil2 := Profil2{
        ID:      1,
        Name:    "Didik Prabowo",
        Age:     40,
        Address: "Wuryantoto, Wonogiri",
        Education: Education{
            SMA: "SMA N 1 Wuryantoro",
            SMP: "SMP N 1 Wuryantoro",
        },
	}
	profil3 := &Profil2{
        ID:      1,
        Name:    "Didik Prabowo",
        Age:     40,
        Address: "Wuryantoto, Wonogiri",
        Education: Education{
            SMA: "SMA N 1 Wuryantoro",
            SMP: "SMP N 1 Wuryantoro",
        },
    }
 
    fmt.Println(profil2)
    fmt.Println("===========================")
    fmt.Printf("Nama \t: %v\n", profil2.Name)
    fmt.Printf("Umur \t: %d\n", profil2.Age)
    fmt.Printf("Alamat \t: %v", profil2.Address)
	fmt.Printf("Sekolah \t: %v", profil2.Education)
	
	
    fmt.Println(*profil3)
    fmt.Println("===========================")
    fmt.Printf("Nama \t: %v\n", (*profil3).Name)
    fmt.Printf("Umur \t: %d\n", (*profil3).Age)
    fmt.Printf("Alamat \t: %v", (*profil3).Address)
    fmt.Printf("Sekolah \t: %v", (*profil3).Education)
}