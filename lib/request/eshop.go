package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ReanSn0w/go-eshop/eshop/models"
	"github.com/ReanSn0w/go-eshop/eshop/request/parameters"
	"github.com/ReanSn0w/go-eshop/eshop/request/region"
)

type EShop struct {
	Items  map[int][]models.Item
	Region region.Region
	Count  int
	Offset int
	Query  string
	Params parameters.Query
}

func (eShop EShop) NextPage() ([]models.Item, error) {
	if eShop.IsLastPage() {
		return nil, fmt.Errorf("по данному запросу уже закончились страницы")
	}

	nextPage := eShop.lastPageNumber() + 1

	requestURL, err := eShop.fullURL()
	if err != nil {
		return nil, err
	}

	nintendoResponse, err := eShop.makeRequest(requestURL)
	if err != nil {
		return nil, err
	}

	eShop.Items[nextPage] = nintendoResponse.Response.Items
	return eShop.Items[nextPage], nil
}

func (eShop EShop) IsLastPage() bool {
	return !((eShop.Offset % eShop.Count) == 0)
}

func (eShop EShop) baseURL() string {
	return fmt.Sprintf("https://searching.nintendo-europe.com/%v/select", eShop.Region)
}

func (eShop EShop) fullURL() (*url.URL, error) {
	values := url.Values{}
	values.Add("q", eShop.Query)
	values.Add("fq", eShop.Params.String())
	values.Add("start", fmt.Sprintf("%v", eShop.lastPageNumber()*eShop.Count))
	values.Add("rows", fmt.Sprintf("%v", eShop.Count))
	values.Add("wt", "json")
	values.Add("bf", "linear(ms(priority,NOW/HOUR),1.1e-11,0)")
	values.Add("bq", "!deprioritise_b:true^999")
	values.Add("json.wrf", "nindo.net.jsonp.jsonpCallback_2772755")
	values.Add("sort", "score desc, date_from desc")

	return url.Parse(fmt.Sprintf("%v?%v", eShop.baseURL(), values.Encode()))
}

func (eShop EShop) makeRequest(url *url.URL) (*models.HTTP, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	webRes := buffer.Bytes()
	webRes = webRes[38 : len(webRes)-2]

	nintendoResponse := models.HTTP{}

	err = json.Unmarshal(webRes, &nintendoResponse)
	return &nintendoResponse, err
}

func (eShop EShop) lastPageNumber() int {
	res := 0

	for page := range eShop.Items {
		if page > res {
			res = page
		}
	}

	return res
}
