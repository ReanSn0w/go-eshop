package request_test

import (
	"log"
	"testing"

	eShop "github.com/ReanSn0w/go-eshop"
	"github.com/ReanSn0w/go-eshop/lib/request/parameters"
	"github.com/ReanSn0w/go-eshop/lib/request/region"
)

func Test_PriceDropQuery(t *testing.T) {
	params := make(parameters.Query)
	params.SetDigitalVersion()
	params.SetPriceDiscount()

	query := eShop.Make("", region.Russia, 10, params)

	items, err := query.NextPage()
	if err != nil {
		t.Error(err)
		return
	}

	for index, item := range items {
		log.Printf("%v) %v\n", index, item.Title)

		if !(item.PriceHasDiscountB && item.DigitalVersionB) {
			t.Error("Данный елемент не соответствует параметрам")
			return
		}
	}
}
