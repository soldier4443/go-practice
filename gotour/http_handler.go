package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, s)
}

type part struct {
	desc string
}

func (p part) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, p.desc)
}

type Robot struct {
	Arm  part
	Body part
	Leg  part
}

func (robot Robot) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	robot.Arm.ServeHTTP(w, r)
	robot.Body.ServeHTTP(w, r)
	robot.Leg.ServeHTTP(w, r)
}

func main() {
	var h Hello

	// Define routes
	http.Handle("/", h)
	http.Handle("/string", String("I'm groot"))
	http.Handle("/robot", Robot{
		part{"These are my arms."},
		part{"This is my body."},
		part{"I have two metal legs."},
	})

	// Use DefaultServeMux to handle the routes
	http.ListenAndServe("localhost:4000", nil)
}
