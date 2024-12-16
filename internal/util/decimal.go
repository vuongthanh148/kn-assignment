package util

import (
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
	"github.com/shopspring/decimal"
)

func NewDecimalFromFloat32(f float32) decimal.Decimal {
	return decimal.NewFromFloat32(f)
}

func NewDecimalFromFloat64(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}

func NewJsonMoneyFromFloat32(f float32) jsonmodel.Money {
	return jsonmodel.Money(NewDecimalFromFloat32(f))
}

func NewDecimalFromUint32(i uint32) decimal.Decimal {
	return decimal.NewFromInt32(int32(i))
}

func NewJsonMoneyFromUint32(i uint32) jsonmodel.Money {
	return jsonmodel.Money(NewDecimalFromUint32(i))
}
