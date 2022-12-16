package limit

type LimitType string

var LimitTypes = struct {
	Annual  LimitType
	Monthly LimitType
}{
	LimitType("ANNUAL"),
	LimitType("MONTHLY"),
}
