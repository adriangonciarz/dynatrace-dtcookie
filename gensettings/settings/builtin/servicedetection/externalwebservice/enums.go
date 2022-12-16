package externalwebservice

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
	Axis       FrameworkType
	Cxf        FrameworkType
	Hessian    FrameworkType
	JaxWsRi    FrameworkType
	Jboss      FrameworkType
	Jersey     FrameworkType
	Progress   FrameworkType
	Resteasy   FrameworkType
	Restlet    FrameworkType
	Spring     FrameworkType
	Tibco      FrameworkType
	Weblogic   FrameworkType
	Webmethods FrameworkType
	Websphere  FrameworkType
	Wink       FrameworkType
}{
	FrameworkType("AXIS"),
	FrameworkType("CXF"),
	FrameworkType("HESSIAN"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("JBOSS"),
	FrameworkType("JERSEY"),
	FrameworkType("PROGRESS"),
	FrameworkType("RESTEASY"),
	FrameworkType("RESTLET"),
	FrameworkType("SPRING"),
	FrameworkType("TIBCO"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("WEBMETHODS"),
	FrameworkType("WEBSPHERE"),
	FrameworkType("WINK"),
}

type TransformationType string

var TransformationTypes = struct {
	After             TransformationType
	Before            TransformationType
	Between           TransformationType
	RemoveCreditCards TransformationType
	RemoveIbans       TransformationType
	RemoveIps         TransformationType
	RemoveNumbers     TransformationType
	ReplaceBetween    TransformationType
	SplitSelect       TransformationType
	TakeSegments      TransformationType
}{
	TransformationType("AFTER"),
	TransformationType("BEFORE"),
	TransformationType("BETWEEN"),
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("REMOVE_IBANS"),
	TransformationType("REMOVE_IPS"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("TAKE_SEGMENTS"),
}
