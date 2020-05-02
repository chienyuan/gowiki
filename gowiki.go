package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
)

// Page struct
type Page struct {
  Title string
  Body []byte
}

func (p *Page)save() error {
  filename := p.Title+ ".txt"
  return ioutil.WriteFile(filename,p.Body,0600)
}

func loadPage(title string) (*Page,error){
  filename := title + ".txt"
  body, err  := ioutil.ReadFile(filename)
  if err != nil {
    return nil,err
  }
  return &Page{Title: title, Body: body},nil
}

func handler(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Hi, I love %s",r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter,r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w,"<h1>%s</h1><div>%s</div",p.Title,string(p.Body))
}

func main(){
  //p1 := &Page{Title: "TestPage", Body: []byte("This is a test page.") }
  //p1.save()
  //p2, _ := loadPage("TestPage")
  //fmt.Println(string(p2.Body))
  http.HandleFunc("/view/",viewHandler)
  fmt.Println("Server up!")
  log.Fatal(http.ListenAndServe(":8080",nil) )
}