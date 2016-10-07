package main

type Links struct {
	Self string `json:"self"`
	Related string `json:"related"`
	Next string `json:"next"`
	Last string `json:"last"`
}

type Resource struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Links *Links `json:"links,omitempty"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}


