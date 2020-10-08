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
