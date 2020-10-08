package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

//Page data structure
type Page struct{
	Title string
	//byte slice
	Body []byte
}

//save method
func (p *Page) save() error{
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".html"
	//underscore used to throw away return value
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil{
		return404(w, r.URL.Path[len("/view/"):])
	}else{
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

func return404(w http.ResponseWriter, req string){
	fmt.Fprintf(w, "<h1>Error 404</h1><div>%s not found</div>",req)
}

func main(){
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

