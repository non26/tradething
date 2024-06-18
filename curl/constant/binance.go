package constant

var Bn_Header_BMX = "X-MBX-APIKEY"
var Bn_Header_Content_Type = "CONTENT-TYPE"
var Bn_Content_Type = "application/x-www-form-urlencoded"

var Bn_CreateOrder = "BN Create Order"
var Bn_QueryOrder = "BN Query Order"
var Bn_SetLeverage = "BN Set Leverage"
var Bn_GetLeverage = "BN Get Leverage"

var BinanceApis = []string{
	Bn_CreateOrder,
	Bn_QueryOrder,
	Bn_SetLeverage,
	Bn_GetLeverage,
}
