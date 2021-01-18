package main
 
import (
    "fmt"
)
 
func main() {
 	for index := 0; index < 10; index++ {
        fmt.Println("Perulangan ke", index)
	}

	var index = 0
 
	fmt.Println("--------- While -----------")
    for index < 4 {
        fmt.Println("Angka ke", index)
        index++
	}
	
	fmt.Println("--------- While Break -----------")

	var indexb = 0

	for indexb < 4 {

		fmt.Println("Angka ke", indexb)
		indexb++

		if indexb == 2 {
			break
		}
	}
	
	fmt.Println("--------- Range -----------")
	
	data := []string{"didik", "Prabowo"}
 
    for index, v := range data {
        fmt.Println("Perulangan ke", index, v)
    }
}