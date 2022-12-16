package customunit

type UnitCombinations string

var UnitCombinationss = struct {
	Scalar   UnitCombinations
	Quotient UnitCombinations
	Product  UnitCombinations
	Power    UnitCombinations
}{
	UnitCombinations("SCALAR"),
	UnitCombinations("QUOTIENT"),
	UnitCombinations("PRODUCT"),
	UnitCombinations("POWER"),
}
