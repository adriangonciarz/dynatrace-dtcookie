package infrastructurehosts

type ConnectionLostDetectionSensitivity string

var ConnectionLostDetectionSensitivities = struct {
	DontAlertOnGracefulShutdown ConnectionLostDetectionSensitivity
	AlertOnGracefulShutdown     ConnectionLostDetectionSensitivity
}{
	ConnectionLostDetectionSensitivity("DONT_ALERT_ON_GRACEFUL_SHUTDOWN"),
	ConnectionLostDetectionSensitivity("ALERT_ON_GRACEFUL_SHUTDOWN"),
}

type DetectionMode string

var DetectionModes = struct {
	Auto   DetectionMode
	Custom DetectionMode
}{
	DetectionMode("Auto"),
	DetectionMode("Custom"),
}
