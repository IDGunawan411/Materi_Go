	
package main
 
import (
    "fmt"
)

const UMR int = 2000000

type Pekerja struct {
    Nama string
    Gaji int
}

func (p Pekerja) lihatPekerja() bool{
    fmt.Println("Nama\t :", p.Nama)
    fmt.Println("Gaji\t :", p.Gaji)

 
	if p.Gaji < UMR {
        return false
    } else {
        return true
	}
}

// ---------------
type (
    Category struct {
        Name string
    }
 
    Post struct {
        Title string
    }
)

func (c Category) lihatData() {
    fmt.Println(c)
}
func (p Post) lihatData() {
    fmt.Println(p)
}

// -------------
type (
    Category2 struct {
        Name string
    }
)
 
func (c Category2) lihatKategori() {
    c.Name = "Berita"
}
func (c *Category2) lihatKategoriLagi() {
    c.Name = "Teknologi"
}

// -----------------
type val int
 
func (i val) jumlah() val {
    s := i * val(2)
    return s
}

 
func main() {
    p1 := Pekerja{
        Nama: "Charly Van Houten",
        Gaji: 1000000,
    }
	
	fmt.Println(p1.lihatPekerja())
	
	fmt.Println("----------------------")
	
	fmt.Printf("From Category\n")
    cats := Category{
        Name: "Berita",
    }

    cats.lihatData()
 
    fmt.Printf("From Post\n")
 
    p := Post{
        Title: "Belajar Golang",
    }
	p.lihatData()

	fmt.Println("----------------------")	
	cats2 := Category2{
        Name: "Info",
    }
 
    cats2.lihatKategori()
    fmt.Println(cats2)
 
    (&cats2).lihatKategoriLagi()
	fmt.Println(cats2)
	
	fmt.Println("----------------------")	
	
    numb := val(10)
    fmt.Println(numb.jumlah())
}