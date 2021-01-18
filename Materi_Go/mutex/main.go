	
package main
 
import (
    "fmt"
    "sync"
)
 
type Data struct {
	ID int
	sy sync.Mutex
}
 
func (s *Data) tambah(o int, wg *sync.WaitGroup) {
	s.sy.Lock()
    s.ID = s.ID + o
    s.sy.Unlock()
    defer wg.Done()
}
 
func main() {
    var dt Data
    var wg sync.WaitGroup
    for i := 0; i < 500; i++ {
        wg.Add(1)
        go func() {
            dt.tambah(1, &wg)
        }()
    }
    wg.Wait()
    fmt.Println(dt.ID)
}