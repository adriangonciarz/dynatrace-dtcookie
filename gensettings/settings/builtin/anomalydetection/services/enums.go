package services

type DetectionMode string

var DetectionModes = struct {
	Auto  DetectionMode
	Fixed DetectionMode
}{
	DetectionMode("Auto"),
	DetectionMode("Fixed"),
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
