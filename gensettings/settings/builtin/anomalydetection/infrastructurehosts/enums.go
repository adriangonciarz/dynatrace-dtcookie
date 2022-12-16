package infrastructurehosts

type ConnectionLostDetectionSensitivity string

var ConnectionLostDetectionSensitivities = struct {
	AlertOnGracefulShutdown     ConnectionLostDetectionSensitivity
	DontAlertOnGracefulShutdown ConnectionLostDetectionSensitivity
}{
	ConnectionLostDetectionSensitivity("ALERT_ON_GRACEFUL_SHUTDOWN"),
	ConnectionLostDetectionSensitivity("DONT_ALERT_ON_GRACEFUL_SHUTDOWN"),
}

type DetectionMode string

var DetectionModes = struct {
	Auto   DetectionMode
	Custom DetectionMode
}{
	DetectionMode("Auto"),
	DetectionMode("Custom"),
}
