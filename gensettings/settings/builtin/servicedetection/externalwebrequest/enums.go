package externalwebrequest

type FrameworkType string

var FrameworkTypes = struct {
	Hessian    FrameworkType
	Websphere  FrameworkType
	Axis       FrameworkType
	Restlet    FrameworkType
	Jersey     FrameworkType
	Progress   FrameworkType
	Wink       FrameworkType
	JaxWsRi    FrameworkType
	Webmethods FrameworkType
	Cxf        FrameworkType
	Weblogic   FrameworkType
	Jboss      FrameworkType
	Resteasy   FrameworkType
	Spring     FrameworkType
	Tibco      FrameworkType
}{
	FrameworkType("HESSIAN"),
	FrameworkType("WEBSPHERE"),
	FrameworkType("AXIS"),
	FrameworkType("RESTLET"),
	FrameworkType("JERSEY"),
	FrameworkType("PROGRESS"),
	FrameworkType("WINK"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("WEBMETHODS"),
	FrameworkType("CXF"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("JBOSS"),
	FrameworkType("RESTEASY"),
	FrameworkType("SPRING"),
	FrameworkType("TIBCO"),
}

type ContextRootTransformationType string

var ContextRootTransformationTypes = struct {
	RemoveCreditCards ContextRootTransformationType
	RemoveIbans       ContextRootTransformationType
	RemoveIps         ContextRootTransformationType
	Before            ContextRootTransformationType
	ReplaceBetween    ContextRootTransformationType
	RemoveNumbers     ContextRootTransformationType
}{
	ContextRootTransformationType("REMOVE_CREDIT_CARDS"),
	ContextRootTransformationType("REMOVE_IBANS"),
	ContextRootTransformationType("REMOVE_IPS"),
	ContextRootTransformationType("BEFORE"),
	ContextRootTransformationType("REPLACE_BETWEEN"),
	ContextRootTransformationType("REMOVE_NUMBERS"),
}

type TransformationType string

var TransformationTypes = struct {
	RemoveCreditCards TransformationType
	RemoveIbans       TransformationType
	TakeSegments      TransformationType
	SplitSelect       TransformationType
	After             TransformationType
	RemoveNumbers     TransformationType
	Before            TransformationType
	Between           TransformationType
	ReplaceBetween    TransformationType
	RemoveIps         TransformationType
}{
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("REMOVE_IBANS"),
	TransformationType("TAKE_SEGMENTS"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("AFTER"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("BEFORE"),
	TransformationType("BETWEEN"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("REMOVE_IPS"),
}

type ContributionType string

var ContributionTypes = struct {
	Originalvalue  ContributionType
	Transformvalue ContributionType
	Transformurl   ContributionType
}{
	ContributionType("OriginalValue"),
	ContributionType("TransformValue"),
	ContributionType("TransformURL"),
}

type ContributionTypeWithOverride string

var ContributionTypeWithOverrides = struct {
	Originalvalue  ContributionTypeWithOverride
	Overridevalue  ContributionTypeWithOverride
	Transformvalue ContributionTypeWithOverride
}{
	ContributionTypeWithOverride("OriginalValue"),
	ContributionTypeWithOverride("OverrideValue"),
	ContributionTypeWithOverride("TransformValue"),
}
