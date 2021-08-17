package models

type Response struct {
	NumFound int64  `json:"numFound"`
	Start    int64  `json:"start"`
	Items    []Item `json:"docs"`
}
