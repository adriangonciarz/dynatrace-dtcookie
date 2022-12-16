package infrastructurevmware

type DetectionMode string

var DetectionModes = struct {
	Auto   DetectionMode
	Custom DetectionMode
}{
	DetectionMode("Auto"),
	DetectionMode("Custom"),
}
