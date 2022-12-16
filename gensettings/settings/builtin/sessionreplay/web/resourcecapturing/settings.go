package resourcecapturing

import "github.com/dtcookie/hcl"

type Settings struct {
	EnableResourceCapturing                bool     `json:"enableResourceCapturing"`                // When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).
	ResourceCaptureUrlExclusionPatternList []string `json:"resourceCaptureUrlExclusionPatternList"` // Add exclusion rules to avoid the capture of resources from certain pages.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enable_resource_capturing": {
			Type:        hcl.TypeBool,
			Description: "When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).",
			Optional:    true,
		},
		"resource_capture_url_exclusion_pattern_list": {
			Type:        hcl.TypeSet,
			Description: "Add exclusion rules to avoid the capture of resources from certain pages.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enable_resource_capturing":                   me.EnableResourceCapturing,
		"resource_capture_url_exclusion_pattern_list": me.ResourceCaptureUrlExclusionPatternList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_resource_capturing":                   &me.EnableResourceCapturing,
		"resource_capture_url_exclusion_pattern_list": &me.ResourceCaptureUrlExclusionPatternList,
	})
}
