package main

import (
	"fmt"
	"tradething/config"
	thisconstant "tradething/curl/constant"
	thisinit "tradething/curl/init"

	"github.com/manifoldco/promptui"
)

func main() {
	config, err := config.ReadConfig("./config")
	if err != nil {
		panic(err.Error())
	}
	// select exchange
	exchanges := thisinit.InitExchangeMenu()[thisconstant.ExchangesKey]
	exchange := promptui.Select{
		Label: thisconstant.ChooseExchanges,
		Items: exchanges,
	}
	_, exchange_result, err := exchange.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n as exchange", err)
		return
	}
	fmt.Printf("You choose %q\n as exchange", exchange_result)

	// select api
	apis := thisinit.InitExchangeApiMenu()[exchange_result]
	api := promptui.Select{
		Label: thisconstant.SelectApi,
		Items: apis,
	}
	_, api_result, err := api.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n as api", err)
		return
	}
	fmt.Printf("You choose %q\n as api", api_result)

	api_maps := thisinit.InitApiRequestBodyMenu()
	select_exchange := api_maps[exchange_result]
	_ = select_exchange
	select_api, err := api_maps[exchange_result][api_result](config)
	if err != nil {
		fmt.Printf("Cann't Get Api %v\n", err)
		return
	}
	field_value := []string{}
	fields := select_api.GetUserInputField()
	for _, field := range fields {
		label := fmt.Sprintf(thisconstant.InputFiledValue, field)
		fieldPrompt := promptui.Prompt{
			Label: label,
		}
		field_result, err := fieldPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n as user input", err)
			return
		}
		field_value = append(field_value, field_result)
		fmt.Printf("You choose %q\n as field value", field_result)
	}

	err = select_api.SetUserInputValue(field_value)
	if err != nil {
		fmt.Printf("Set User Input Fail %v\n", err)
	}

	err = select_api.ExecuteCurl()
	if err != nil {
		fmt.Printf("Execute Curl Fail %v\n", err)
	}

}
