package maintenance

import "github.com/dtcookie/hcl"

type GeneralProperties struct {
	Name             string      `json:"name"`                             // The name of the maintenance window, displayed in the UI.
	Description      *string     `json:"description,omitempty"`            // A short description of the maintenance purpose.
	Type             WindowType  `json:"maintenanceType"`                  // The type of the maintenance: planned or unplanned
	Suppression      Suppression `json:"suppression"`                      // The type of suppression of alerting and problem detection during the maintenance.
	DisableSynthetic bool        `json:"disableSyntheticMonitorExecution"` // Suppress execution of synthetic monitors during the maintenance
}

func (me *GeneralProperties) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the maintenance window, displayed in the UI",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the maintenance purpose",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the maintenance: planned or unplanned",
			Required:    true,
		},
		"suppression": {
			Type:        hcl.TypeString,
			Description: "The type of suppression of alerting and problem detection during the maintenance",
			Required:    true,
		},
		"disable_synthetic": {
			Type:        hcl.TypeBool,
			Description: "Suppress execution of synthetic monitors during the maintenance",
			Default:     false,
			Optional:    true,
		},
	}
}

func (me *GeneralProperties) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"name":              me.Name,
		"description":       me.Description,
		"type":              me.Type,
		"suppression":       me.Suppression,
		"disable_synthetic": me.DisableSynthetic,
	})
}

func (me *GeneralProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":              &me.Name,
		"description":       &me.Description,
		"type":              &me.Type,
		"suppression":       &me.Suppression,
		"disable_synthetic": &me.DisableSynthetic,
	})
}

// WindowType The type of the maintenance: planned or unplanned.
type WindowType string

// MaintenanceWindowTypes offers the known enum values
var MaintenanceWindowTypes = struct {
	Planned   WindowType
	Unplanned WindowType
}{
	"PLANNED",
	"UNPLANNED",
}

// Suppression The type of suppression of alerting and problem detection during the maintenance.
type Suppression string

// Suppressions offers the known enum values
var Suppressions = struct {
	DetectProblemsAndAlert  Suppression
	DetectProblemsDontAlert Suppression
	DontDetectProblems      Suppression
}{
	"DETECT_PROBLEMS_AND_ALERT",
	"DETECT_PROBLEMS_DONT_ALERT",
	"DONT_DETECT_PROBLEMS",
}
