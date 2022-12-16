package generictype

import (
	"github.com/dtcookie/hcl"
)

type ExtractionRules []*ExtractionRule

func (me *ExtractionRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ExtractionRule).Schema()},
		},
	}
}

func (me ExtractionRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("rule", me)
}

func (me *ExtractionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("rule", me)
}

// Entity extraction rule. An extraction rule defines which sources are evaluated for extracting an entity. If a source matches the specified filters, an entity is extracted.
type ExtractionRule struct {
	Attributes          AttributeEntries `json:"attributes"`          // All attribute extraction rules will be applied and found attributes will be added to the extracted type.
	IconPattern         string           `json:"iconPattern"`         // Define a pattern which is used to set the icon attribute of the entity. The extracted values must reference barista icon ids. You may define placeholders referencing data source dimensions.
	IdPattern           string           `json:"idPattern"`           // ID patterns are comprised of static text and placeholders referring to dimensions in the ingest data. An ID pattern **must** contain at least one placeholder to ensure that different entities will be created.. Take care that the pattern results in the same ID for the same entity. For example, using timestamp or counter-like dimensions as part of the ID would lead to the creation of new entities for each ingest data and is strongly discouraged!\n\nEach dimension key referred to by an identifier placeholder must be present in order to extract an entity. If any dimension key referred to in the identifier is missing, the rule will not be considered for evaluation. If you have cases where you still want to extract the same entity type but have differently named keys, consider creating multiple rules extracting the same entity type. In this case take care that each ID pattern evaluates to the same value if the same entity should be extracted.
	InstanceNamePattern string           `json:"instanceNamePattern"` // Define a pattern which is used to set the name attribute of the entity. You may define placeholders referencing data source dimensions.
	RequiredDimensions  DimensionFilters `json:"requiredDimensions"`  // In addition to the dimensions already referred to in the ID pattern, you may specify additional dimensions which must be present in order to evaluate this rule.
	Role                string           `json:"role"`                // If you want to extract multiple entities of the same type from a single ingest line you need to define multiple rules with different roles.
	Sources             SourceFilters    `json:"sources"`             // Specify all sources which should be evaluated for this rule. A rule is evaluated if any of the specified source filters match.
}

func (me *ExtractionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attributes": {
			Type:        hcl.TypeList,
			Description: "All attribute extraction rules will be applied and found attributes will be added to the extracted type.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AttributeEntries).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"icon_pattern": {
			Type:        hcl.TypeString,
			Description: "Define a pattern which is used to set the icon attribute of the entity. The extracted values must reference barista icon ids. You may define placeholders referencing data source dimensions.",
			Optional:    true,
		},
		"id_pattern": {
			Type:        hcl.TypeString,
			Description: "ID patterns are comprised of static text and placeholders referring to dimensions in the ingest data. An ID pattern **must** contain at least one placeholder to ensure that different entities will be created.. Take care that the pattern results in the same ID for the same entity. For example, using timestamp or counter-like dimensions as part of the ID would lead to the creation of new entities for each ingest data and is strongly discouraged!\n\nEach dimension key referred to by an identifier placeholder must be present in order to extract an entity. If any dimension key referred to in the identifier is missing, the rule will not be considered for evaluation. If you have cases where you still want to extract the same entity type but have differently named keys, consider creating multiple rules extracting the same entity type. In this case take care that each ID pattern evaluates to the same value if the same entity should be extracted.",
			Required:    true,
		},
		"instance_name_pattern": {
			Type:        hcl.TypeString,
			Description: "Define a pattern which is used to set the name attribute of the entity. You may define placeholders referencing data source dimensions.",
			Optional:    true,
		},
		"required_dimensions": {
			Type:        hcl.TypeList,
			Description: "In addition to the dimensions already referred to in the ID pattern, you may specify additional dimensions which must be present in order to evaluate this rule.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DimensionFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"role": {
			Type:        hcl.TypeString,
			Description: "If you want to extract multiple entities of the same type from a single ingest line you need to define multiple rules with different roles.",
			Optional:    true,
		},
		"sources": {
			Type:        hcl.TypeList,
			Description: "Specify all sources which should be evaluated for this rule. A rule is evaluated if any of the specified source filters match.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SourceFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ExtractionRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attributes":            me.Attributes,
		"icon_pattern":          me.IconPattern,
		"id_pattern":            me.IdPattern,
		"instance_name_pattern": me.InstanceNamePattern,
		"required_dimensions":   me.RequiredDimensions,
		"role":                  me.Role,
		"sources":               me.Sources,
	})
}

func (me *ExtractionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attributes":            &me.Attributes,
		"icon_pattern":          &me.IconPattern,
		"id_pattern":            &me.IdPattern,
		"instance_name_pattern": &me.InstanceNamePattern,
		"required_dimensions":   &me.RequiredDimensions,
		"role":                  &me.Role,
		"sources":               &me.Sources,
	})
}
