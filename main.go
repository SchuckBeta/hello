package main

import (
        "fmt"
        "log"
        "net/http"
        "io/ioutil"
)


type Page struct {
    Title string
    Body  []byte
}


func HelloWorld(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!!!"))
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}


func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}


func main() {
        p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
        p1.save()
        p2, _ := loadPage("TestPage")
        fmt.Println(string(p2.Body))


        fmt.Println("hello world")
        http.HandleFunc("/", HelloWorld)
        log.Fatal(http.ListenAndServe(":80", nil))
}
