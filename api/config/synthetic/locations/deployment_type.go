package locations

type DeploymentType string

var DeploymentTypes = struct {
	Kubernetes DeploymentType
	Standard   DeploymentType
}{
	DeploymentType("KUBERNETES"),
	DeploymentType("STANDARD"),
}

func (me DeploymentType) Ref() *DeploymentType {
	return &me
}
