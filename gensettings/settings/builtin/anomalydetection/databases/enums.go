package databases

type DetectionMode string

var DetectionModes = struct {
	Fixed DetectionMode
	Auto  DetectionMode
}{
	DetectionMode("Fixed"),
	DetectionMode("Auto"),
}

type Sensitivity string

var Sensitivities = struct {
	High   Sensitivity
	Low    Sensitivity
	Medium Sensitivity
}{
	Sensitivity("High"),
	Sensitivity("Low"),
	Sensitivity("Medium"),
}
