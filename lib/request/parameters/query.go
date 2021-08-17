package parameters

import "fmt"

type Query map[string][]string

func (params Query) String() string {
	return fmt.Sprintf("type:GAME AND (%v) AND sorting_title:* AND *:*", params.internalParams())
}

func (params Query) SetPhysicalVersion() {
	params.setRequestParameter("physical_version_b", "true", true)
}

func (params Query) SetDigitalVersion() {
	params.setRequestParameter("digital_version_b", "true", true)
}

func (params Query) SetPriceDiscount() {
	params.setRequestParameter("price_has_discount_b", "true", true)
}

func (params Query) SetDemoAvalability() {
	params.setRequestParameter("demo_availability", "true", true)
}

func (params Query) SetDLCAvalability() {
	params.setRequestParameter("add_on_content_b", "true", true)
}

func (params Query) setRequestParameter(key string, value string, force bool) {
	if params[key] == nil || force {
		params[key] = []string{value}
	} else {
		for _, item := range params[key] {
			if value == item {
				return
			}
		}

		params[key] = append(params[key], value)
	}
}

func (params Query) internalParams() string {
	res := ""

	for key, values := range params {
		if res != "" {
			res += " AND "
		}
		res += params.renderValues(key, values)
	}

	return fmt.Sprintf("(%v)", res)
}

func (params Query) renderValues(key string, values []string) string {
	res := ""

	for index, value := range values {
		if index == 0 {
			res = fmt.Sprintf("%v:\"%v\"", key, value)
			continue
		}

		res += " OR "
		res += fmt.Sprintf("%v:\"%v\"", key, value)
	}

	return fmt.Sprintf("(%v)", res)
}
