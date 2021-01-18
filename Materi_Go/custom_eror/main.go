package main
import (
    "fmt"
)
type (
	customError struct { 
		err string 
		val int
	}
)
func (ce *customError) Error() string {
    return fmt.Sprintf("Nilainya %d: %s", ce.val, ce.err)
}
func checkAge(nilai int) (int, error) {
    if nilai <= 10 {
        return nilai, &customError{err: "Umur tidak di ijinkan", val: nilai}
    }
 	return nilai, nil
}
func main() {
    checkAge, err := checkAge(10)
    if err != nil {
        if err, val := err.(*customError); val {
            fmt.Printf("Umur anda tidak di ijinkan %d %s\n", err.val, "Sabar")
            return
        }
        fmt.Println(err)
        return
	}
	fmt.Printf("Umur anda dalah %d \n", checkAge)
}