package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
		Bio  string
	}{
		Name: "Samuel Sousa",
		Bio:  "<alert>eu sou eu</alert>",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
