package format

import (
	"github.com/leekchan/accounting"
)

// ConvertInt4ToCurrency ...
func ConvertInt4ToCurrency(value int64) string {
	ac := accounting.Accounting{Thousand: "."}
	return ac.FormatMoney(value)
}

// ConvertFloat64ToCurrency ...
func ConvertFloat64ToCurrency(value float64) string {
	ac := accounting.Accounting{Thousand: "."}
	return ac.FormatMoney(value)
}

// ToCurrencyVND ...
func ToCurrencyVND(value interface{}) string {
	ac := accounting.Accounting{Symbol: "", Precision: 0, Thousand: "."}
	symbol := "đ"
	return ac.FormatMoney(value) + symbol
}
