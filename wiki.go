package main

import (
	"fmt"
	"io/ioutil"
)

//Page data structure
type Page struct{
	Title string
	//byte slice
	Body []byte
}

//save method
func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page{
	filename := title + ".txt"
	//underscore used to throw away return value
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

