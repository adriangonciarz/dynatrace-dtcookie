package rumcustom

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
	Low    Sensitivity
	Medium Sensitivity
	High   Sensitivity
}{
	Sensitivity("Low"),
	Sensitivity("Medium"),
	Sensitivity("High"),
}
