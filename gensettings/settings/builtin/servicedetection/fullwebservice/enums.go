package fullwebservice

type TransformationType string

var TransformationTypes = struct {
	Before            TransformationType
	RemoveCreditCards TransformationType
	RemoveNumbers     TransformationType
	RemoveIps         TransformationType
	SplitSelect       TransformationType
	After             TransformationType
	TakeSegments      TransformationType
	Between           TransformationType
	ReplaceBetween    TransformationType
	RemoveIbans       TransformationType
}{
	TransformationType("BEFORE"),
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("REMOVE_IPS"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("AFTER"),
	TransformationType("TAKE_SEGMENTS"),
	TransformationType("BETWEEN"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("REMOVE_IBANS"),
}

type ContributionType string

var ContributionTypes = struct {
	Transformvalue ContributionType
	Transformurl   ContributionType
	Originalvalue  ContributionType
}{
	ContributionType("TransformValue"),
	ContributionType("TransformURL"),
	ContributionType("OriginalValue"),
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

type FrameworkType string

var FrameworkTypes = struct {
	Hessian    FrameworkType
	Spring     FrameworkType
	Webmethods FrameworkType
	Axis       FrameworkType
	Jboss      FrameworkType
	Restlet    FrameworkType
	Jersey     FrameworkType
	Cxf        FrameworkType
	JaxWsRi    FrameworkType
	Tibco      FrameworkType
	Progress   FrameworkType
	Weblogic   FrameworkType
	Wink       FrameworkType
	Resteasy   FrameworkType
	Websphere  FrameworkType
}{
	FrameworkType("HESSIAN"),
	FrameworkType("SPRING"),
	FrameworkType("WEBMETHODS"),
	FrameworkType("AXIS"),
	FrameworkType("JBOSS"),
	FrameworkType("RESTLET"),
	FrameworkType("JERSEY"),
	FrameworkType("CXF"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("TIBCO"),
	FrameworkType("PROGRESS"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("WINK"),
	FrameworkType("RESTEASY"),
	FrameworkType("WEBSPHERE"),
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
