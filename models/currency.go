package models

type Currency string

const (
	USD Currency = "USD" // amaerican dollar
	EUR Currency = "EUR" // euro
	COP Currency = "COP" // colombian peso
	CAD Currency = "CAD" // canadian dollar
	MXN Currency = "MXN" // mexican peso
	JPY Currency = "JPY" // japanse yen
	CNY Currency = "CNY" // china yuan
)

func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case COP:
		return "COP"
	case CAD:
		return "CAD"
	case MXN:
		return "MXN"
	case JPY:
		return "JPY"
	case CNY:
		return "CNY"
	}
	return "unknown"
}

func SupportedCurrencies() []string {
	var supported = []string{
		USD.String(),
		EUR.String(),
		COP.String(),
		CAD.String(),
		MXN.String(),
		JPY.String(),
		CNY.String(),
	}
	return supported
}
