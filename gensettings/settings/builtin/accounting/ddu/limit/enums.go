package limit

type LimitType string

var LimitTypes = struct {
	Monthly LimitType
	Annual  LimitType
}{
	LimitType("MONTHLY"),
	LimitType("ANNUAL"),
}
