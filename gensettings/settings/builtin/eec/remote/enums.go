package remote

type PerformanceProfile string

var PerformanceProfiles = struct {
	High    PerformanceProfile
	Default PerformanceProfile
}{
	PerformanceProfile("HIGH"),
	PerformanceProfile("DEFAULT"),
}
