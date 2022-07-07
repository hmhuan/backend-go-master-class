package util

var currencies = map[string]string{
	"USD": "United States Dollar",
	"EUR": "Euro",
	"VND": "Vietname Dong",
}

func isSupportedCurrency(currency string) bool {
	_, supported := currencies[currency]
	return supported
}
