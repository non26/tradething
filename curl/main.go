package main

import (
	"fmt"
	"tradething/config"
	thisconstant "tradething/curl/constant"
	thisinit "tradething/curl/init"

	"github.com/manifoldco/promptui"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}
	// select exchange
	exchange := promptui.Select{
		Label: thisconstant.ChooseExchanges,
		Items: thisinit.InitExchangeMenu()[thisconstant.ExchangesKey],
	}
	_, exchangeResult, err := exchange.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n as exchange", err)
		return
	}
	fmt.Printf("You choose %q\n as exchange", exchangeResult)

	// select api
	api := promptui.Select{
		Label: thisconstant.SelectApi,
		Items: thisinit.InitExchangeApiMenu()[exchangeResult],
	}
	_, apiResult, err := api.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n as api", err)
		return
	}
	fmt.Printf("You choose %q\n as api", apiResult)
	apiMap := thisinit.InitApiRequestBodyMenu()

	newApi, err := apiMap[exchangeResult][apiResult](config)
	if err != nil {
		fmt.Printf("Cann't Get Api %v\n", err)
		return
	}
	fieldValue := []string{}
	fields := newApi.GetUserInputField()
	for _, field := range fields {
		label := fmt.Sprintf(thisconstant.InputFiledValue, field)
		fieldPrompt := promptui.Prompt{
			Label: label,
		}
		fieldResult, err := fieldPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n as user input", err)
			return
		}
		fieldValue = append(fieldValue, fieldResult)
		fmt.Printf("You choose %q\n as field value", fieldResult)
	}

	err = newApi.SetUserInputValue(fieldValue)
	if err != nil {
		fmt.Printf("Set User Input Fail %v\n", err)
	}

	err = newApi.GenerateCurl()
	if err != nil {
		fmt.Printf("Execute Curl Fail %v\n", err)
	}

}
