package locations

type CloudPlatform string

var CloudPlatforms = struct {
	Alibaba        CloudPlatform
	AmazonEC2      CloudPlatform
	Azure          CloudPlatform
	DynatraceCloud CloudPlatform
	GoogleCloud    CloudPlatform
	Interoute      CloudPlatform
	Other          CloudPlatform
	Undefined      CloudPlatform
}{
	CloudPlatform("ALIBABA"),
	CloudPlatform("AMAZON_EC2"),
	CloudPlatform("AZURE"),
	CloudPlatform("DYNATRACE_CLOUD"),
	CloudPlatform("GOOGLE_CLOUD"),
	CloudPlatform("INTEROUTE"),
	CloudPlatform("OTHER"),
	CloudPlatform("UNDEFINED"),
}
