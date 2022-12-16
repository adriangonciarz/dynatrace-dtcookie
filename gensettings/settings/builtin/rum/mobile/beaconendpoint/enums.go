package beaconendpoint

type BeaconEndpointType string

var BeaconEndpointTypes = struct {
	InstrumentedWebserver BeaconEndpointType
	ClusterActivegate     BeaconEndpointType
	EnvironmentActivegate BeaconEndpointType
}{
	BeaconEndpointType("INSTRUMENTED_WEBSERVER"),
	BeaconEndpointType("CLUSTER_ACTIVEGATE"),
	BeaconEndpointType("ENVIRONMENT_ACTIVEGATE"),
}
