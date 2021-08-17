package models

type HTTP struct {
	Headers  Header   `json:"responseHeader"`
	Response Response `json:"response"`
}
