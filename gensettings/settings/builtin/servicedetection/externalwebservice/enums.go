package externalwebservice

type TransformationType string

var TransformationTypes = struct {
	Before            TransformationType
	After             TransformationType
	RemoveCreditCards TransformationType
	SplitSelect       TransformationType
	RemoveNumbers     TransformationType
	ReplaceBetween    TransformationType
	TakeSegments      TransformationType
	Between           TransformationType
	RemoveIbans       TransformationType
	RemoveIps         TransformationType
}{
	TransformationType("BEFORE"),
	TransformationType("AFTER"),
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("TAKE_SEGMENTS"),
	TransformationType("BETWEEN"),
	TransformationType("REMOVE_IBANS"),
	TransformationType("REMOVE_IPS"),
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

type FrameworkType string

var FrameworkTypes = struct {
	Websphere  FrameworkType
	Wink       FrameworkType
	Spring     FrameworkType
	Jersey     FrameworkType
	Restlet    FrameworkType
	Tibco      FrameworkType
	Jboss      FrameworkType
	Resteasy   FrameworkType
	Axis       FrameworkType
	JaxWsRi    FrameworkType
	Progress   FrameworkType
	Hessian    FrameworkType
	Cxf        FrameworkType
	Weblogic   FrameworkType
	Webmethods FrameworkType
}{
	FrameworkType("WEBSPHERE"),
	FrameworkType("WINK"),
	FrameworkType("SPRING"),
	FrameworkType("JERSEY"),
	FrameworkType("RESTLET"),
	FrameworkType("TIBCO"),
	FrameworkType("JBOSS"),
	FrameworkType("RESTEASY"),
	FrameworkType("AXIS"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("PROGRESS"),
	FrameworkType("HESSIAN"),
	FrameworkType("CXF"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("WEBMETHODS"),
}
