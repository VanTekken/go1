package main

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"html/template"
)

//In-Memory Storage
type Page struct{
	Title string
	Body []byte
}

//Persistent Storage
func (p *Page) save() error{
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error){
    filename := "templates/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
//Render html template file
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles("templates/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

//MVC Router
func homeHandler(w http.ResponseWriter, r *http.Request) {
    //Redirect root dir requests to the view page
    http.Redirect(w, r, "/view/",http.StatusFound)
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if p == nil{
        renderTemplate(w, "welcome", p)
    }else if err != nil{
        http.Redirect(w, r, "/error/",http.StatusFound)
        return
    }else{
        renderTemplate(w, "view", p)
    }
    
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    if title == "" {
        renderTemplate(w, "edit-help", p)
    }else{
        renderTemplate(w, "edit", p)
    }
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main(){
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/welcome", viewHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}