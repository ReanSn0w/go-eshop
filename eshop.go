package request

import (
	"github.com/ReanSn0w/go-eshop/lib/models"
	"github.com/ReanSn0w/go-eshop/lib/request"
	"github.com/ReanSn0w/go-eshop/lib/request/parameters"
	"github.com/ReanSn0w/go-eshop/lib/request/region"
)

func Make(query string, region region.Region, count int, params parameters.Query) request.EShop {
	return request.EShop{
		Items:  make(map[int][]models.Item),
		Region: region,
		Count:  count,
		Offset: 0,
		Query: func() string {
			if len(query) == 0 {
				return "*"
			} else {
				return query
			}
		}(),
		Params: params,
	}
}
