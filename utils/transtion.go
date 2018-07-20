package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func TransYuanToLi(amount interface{}) (int64, error) {
	return transAnyToLi(amount, 1000)
}

func TransFenToLi(amount interface{}) (int64, error) {
	return transAnyToLi(amount, 10)
}

func TransLiToFen(amount interface{}) (float64, error) {
	return transLiToAny(amount, 10)
}

func TransLiToYuan(amount interface{}) (float64, error) {
	return transLiToAny(amount, 1000)
}

func transAnyToLi(amount interface{}, rate float64) (int64, error) {
	var (
		fee float64
		err error
	)
	switch v := amount.(type) {
	case string:
		fee, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
	case int64:
		fee = float64(v)
	case float64:
		fee = v
	}

	money := decimal.NewFromFloat(fee).Mul(decimal.NewFromFloat(rate)).String()
	rsMoney, err := strconv.ParseInt(money, 10, 64)
	if err != nil {
		return 0, err
	}

	return rsMoney, nil
}

func transLiToAny(amount interface{}, rate float64) (float64, error) {
	var (
		fee float64
		err error
	)

	switch v := amount.(type) {
	case string:
		fee, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
	case int64:
		fee = float64(v)
	case float64:
		fee = v
	}

	money := decimal.NewFromFloat(fee).Div(decimal.NewFromFloat(rate)).String()
	rsMoney, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return 0, err
	}

	return rsMoney, nil
}
