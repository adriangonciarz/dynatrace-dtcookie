package fullwebrequest

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

type TransformationType string

var TransformationTypes = struct {
	RemoveIbans       TransformationType
	TakeSegments      TransformationType
	RemoveIps         TransformationType
	SplitSelect       TransformationType
	ReplaceBetween    TransformationType
	RemoveNumbers     TransformationType
	RemoveCreditCards TransformationType
	Before            TransformationType
	After             TransformationType
	Between           TransformationType
}{
	TransformationType("REMOVE_IBANS"),
	TransformationType("TAKE_SEGMENTS"),
	TransformationType("REMOVE_IPS"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("BEFORE"),
	TransformationType("AFTER"),
	TransformationType("BETWEEN"),
}

type FrameworkType string

var FrameworkTypes = struct {
	Jersey     FrameworkType
	Tibco      FrameworkType
	Cxf        FrameworkType
	Resteasy   FrameworkType
	Hessian    FrameworkType
	Restlet    FrameworkType
	JaxWsRi    FrameworkType
	Weblogic   FrameworkType
	Jboss      FrameworkType
	Progress   FrameworkType
	Spring     FrameworkType
	Wink       FrameworkType
	Websphere  FrameworkType
	Axis       FrameworkType
	Webmethods FrameworkType
}{
	FrameworkType("JERSEY"),
	FrameworkType("TIBCO"),
	FrameworkType("CXF"),
	FrameworkType("RESTEASY"),
	FrameworkType("HESSIAN"),
	FrameworkType("RESTLET"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("JBOSS"),
	FrameworkType("PROGRESS"),
	FrameworkType("SPRING"),
	FrameworkType("WINK"),
	FrameworkType("WEBSPHERE"),
	FrameworkType("AXIS"),
	FrameworkType("WEBMETHODS"),
}

type ContextRootTransformationType string

var ContextRootTransformationTypes = struct {
	RemoveIps         ContextRootTransformationType
	Before            ContextRootTransformationType
	ReplaceBetween    ContextRootTransformationType
	RemoveNumbers     ContextRootTransformationType
	RemoveCreditCards ContextRootTransformationType
	RemoveIbans       ContextRootTransformationType
}{
	ContextRootTransformationType("REMOVE_IPS"),
	ContextRootTransformationType("BEFORE"),
	ContextRootTransformationType("REPLACE_BETWEEN"),
	ContextRootTransformationType("REMOVE_NUMBERS"),
	ContextRootTransformationType("REMOVE_CREDIT_CARDS"),
	ContextRootTransformationType("REMOVE_IBANS"),
}
