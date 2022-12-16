package rumcustomcrashrateincrease

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

type DetectionMode string

var DetectionModes = struct {
	Auto  DetectionMode
	Fixed DetectionMode
}{
	DetectionMode("Auto"),
	DetectionMode("Fixed"),
}
