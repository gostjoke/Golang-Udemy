package furtherexplored

var AValue = 100
var BValue = 200

func ValueExchange() (int, int) { // visible outside the package due to capitalization
	return AValue, BValue
}

func valueExchange() (int, int) { // not visible outside the package
	return AValue, BValue
}
