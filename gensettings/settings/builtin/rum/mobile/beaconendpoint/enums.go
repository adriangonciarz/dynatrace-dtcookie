package beaconendpoint

type BeaconEndpointType string

var BeaconEndpointTypes = struct {
	ClusterActivegate     BeaconEndpointType
	EnvironmentActivegate BeaconEndpointType
	InstrumentedWebserver BeaconEndpointType
}{
	BeaconEndpointType("CLUSTER_ACTIVEGATE"),
	BeaconEndpointType("ENVIRONMENT_ACTIVEGATE"),
	BeaconEndpointType("INSTRUMENTED_WEBSERVER"),
}
