package customunit

type UnitCombinations string

var UnitCombinationss = struct {
	Power    UnitCombinations
	Product  UnitCombinations
	Quotient UnitCombinations
	Scalar   UnitCombinations
}{
	UnitCombinations("POWER"),
	UnitCombinations("PRODUCT"),
	UnitCombinations("QUOTIENT"),
	UnitCombinations("SCALAR"),
}
