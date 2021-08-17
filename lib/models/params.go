package models

type Params struct {
	Q       string `json:"q"`
	JSONWrf string `json:"json.wrf"`
	Bf      string `json:"bf"`
	Start   string `json:"start"`
	Fq      string `json:"fq"`
	Sort    string `json:"sort"`
	Rows    string `json:"rows"`
	Wt      string `json:"wt"`
	Bq      string `json:"bq"`
}
