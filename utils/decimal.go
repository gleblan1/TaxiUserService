package utils

import "github.com/shopspring/decimal"

func ConvertDecimalToInt(value decimal.Decimal) int64 {
	return value.Mul(decimal.NewFromFloat(100)).IntPart()
}

func ConvertFloatToInt(value float64) int64 {
	return decimal.NewFromFloatWithExponent(value, -2).Mul(decimal.NewFromFloat(100)).IntPart()
}

func ConvertIntToFloat(value int64) float64 {
	result, _ := decimal.NewFromFloatWithExponent(float64(value), -2).Div(decimal.NewFromFloat(100)).Float64()
	return result
}
