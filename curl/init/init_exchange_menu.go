package init

import "tradething/curl/constant"

// func InitExchangeMenu() []menuitem.IExchangeMenuItem {
// 	exchanges := []menuitem.IExchangeMenuItem{}
// 	for idx, name := range constant.Exchange {
// 		item := menuitem.NewExchangeMenu(name, idx)
// 		exchanges = append(exchanges, item)
// 	}
// 	return exchanges
// }

func InitExchangeMenu() map[string][]string {
	m := map[string][]string{}
	m[constant.ExchangesKey] = constant.Exchanges
	return m
}
