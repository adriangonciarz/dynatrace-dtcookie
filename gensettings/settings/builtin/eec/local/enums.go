package local

type PerformanceProfile string

var PerformanceProfiles = struct {
	Default PerformanceProfile
	High    PerformanceProfile
}{
	PerformanceProfile("DEFAULT"),
	PerformanceProfile("HIGH"),
}
