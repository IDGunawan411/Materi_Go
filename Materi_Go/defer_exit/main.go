package main
import (
    "fmt"
    "os"
    "sync"
)
var (
    wg sync.WaitGroup
    name string
)
func cetak(name string) {
    defer wg.Done()
    fmt.Println("Nama Saya adalah", name)
}
func main() {
    name := "Didik Prabowo"
    fmt.Println("Halo semuanya...")
    os.Exit(1)
    wg.Add(1)
    go cetak(name)
    wg.Wait()
}