package models

type Header struct {
	Status int64  `json:"status"`
	QTime  int64  `json:"QTime"`
	Params Params `json:"params"`
}
