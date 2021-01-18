package main

import "encoding/json"
import "net/http"
import "fmt"

type student struct {
    ID    string
    Name  string
    Grade int
    Alamat string
}
type mahasiswa struct {
    Nim      string
    Nama     string
    Alamat   string
    Semester string
}

var data = []student{
    student{"E001", "ethan", 21, "Jakarta"},
    student{"W001", "wick", 22 , "Bekasi"},
    student{"B001", "bourne", 23, "Bogor"},
    student{"B002", "bond", 23, "Cibinong"},
}

var data2 = []mahasiswa{
    mahasiswa{"180441180015", "Gunawan","Cibinong","5"},
    mahasiswa{"180441180016", "Gunawan Prasetyo", "Jakarta","4"},
    mahasiswa{"180441180017", "Hendrik", "Jakarta","3"},
    mahasiswa{"180441180016", "Agung", "Bogor","2"},
    mahasiswa{"180441180017", "Hendrik", "Jakarta","4"},
}
func mhs_src(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var sm = r.FormValue("smester")
        var nama = r.FormValue("nama")
        var result []byte
        var err error

        for _, each := range data2 {
            if each.Nama == nama {
                result, err = json.Marshal(each)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                w.Write(result)
                return
            }
            if each.Semester == sm {
                result, err = json.Marshal(each)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func mhs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
        var result, err = json.Marshal(data2)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}
func user(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var id = r.FormValue("id")
        var result []byte
        var err error

        for _, each := range data {
            if each.ID == id {
                result, err = json.Marshal(each)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}
func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func main() {
    http.HandleFunc("/users", users)
    http.HandleFunc("/user", user)
    http.HandleFunc("/mhs", mhs)
    http.HandleFunc("/mhs_src", mhs_src)

    fmt.Println("starting web server at http://localhost:808/")
    http.ListenAndServe(":808", nil)
}