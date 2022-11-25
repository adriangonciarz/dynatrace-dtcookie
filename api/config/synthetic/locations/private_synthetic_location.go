package locations

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type PrivateSyntheticLocation struct {
	ID                               *string         `json:"entityId,omitempty"`                         // The Dynatrace entity ID of the location
	Type                             LocationType    `json:"type"`                                       // Needs to be `PRIVATE`
	Name                             string          `json:"name"`                                       // The name of the location
	CountryCode                      *string         `json:"countryCode,omitempty"`                      // The country code of the location. \n\n Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland)
	RegionCode                       *string         `json:"regionCode,omitempty"`                       // The region code of the location. \n\n For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes (without `US-` or `CA-` prefix), for example, `VA` for Virginia or `OR` for Oregon. \n\n For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes)
	City                             *string         `json:"city,omitempty"`                             // The city of the location
	Latitude                         float64         `json:"latitude"`                                   // The latitude of the location in `DDD.dddd` format
	Longitude                        float64         `json:"longitude"`                                  // The longitude of the location in `DDD.dddd` format
	Status                           *Status         `json:"status,omitempty"`                           // The status of the location: \n\n* `ENABLED`: The location is displayed as active in the UI. You can assign monitors to the location. \n* `DISABLED`: The location is displayed as inactive in the UI. You can't assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location. \n* `HIDDEN`: The location is not displayed in the UI. You can't assign monitors to the location. You can only set location as `HIDDEN` when no monitor is assigned to it
	Nodes                            []string        `json:"nodes,omitempty"`                            // A list of synthetic nodes belonging to the location. \n\n You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call
	AvailabilityLocationOutage       bool            `json:"availabilityLocationOutage,omitempty"`       // The alerting of location outage is enabled (`true`) or disabled (`false`)
	AvailabilityNodeOutage           bool            `json:"availabilityNodeOutage,omitempty"`           // The alerting of node outage is enabled (`true`) or disabled (`false`). \n\n If enabled, the outage of *any* node in the location triggers an alert
	LocationNodeOutageDelayInMinutes *int            `json:"locationNodeOutageDelayInMinutes,omitempty"` // Alert if the location or node outage lasts longer than *X* minutes. \n\n Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to `true`
	AvailabilityNotificationsEnabled bool            `json:"availabilityNotificationsEnabled,omitempty"` // The notifications of location and node outage is enabled (`true`) or disabled (`false`)
	DeploymentType                   *DeploymentType `json:"deploymentType,omitempty"`                   // The deployment type of the location: \n\n* `STANDARD`: The location is deployed on Windows or Linux.\n* `KUBERNETES`: The location is deployed on Kubernetes
	AutoUpdateChromium               bool            `json:"autoUpdateChromium,omitempty"`               // Auto upgrade of Chromium is enabled (`true`) or disabled (`false`)
}

func (me *PrivateSyntheticLocation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the location",
			Required:    true,
		},
		"country_code": {
			Type:        hcl.TypeString,
			Description: "The country code of the location. \n\n Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland)",
			Optional:    true,
		},
		"region_code": {
			Type:        hcl.TypeString,
			Description: "The region code of the location. \n\n For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes (without `US-` or `CA-` prefix), for example, `VA` for Virginia or `OR` for Oregon. \n\n For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes)",
			Optional:    true,
		},
		"city": {
			Type:        hcl.TypeString,
			Description: "The city of the location",
			Optional:    true,
		},
		"latitude": {
			Type:        hcl.TypeFloat,
			Description: "The latitude of the location in `DDD.dddd` format",
			Required:    true,
		},
		"longitude": {
			Type:        hcl.TypeFloat,
			Description: "The longitude of the location in `DDD.dddd` format",
			Required:    true,
		},
		"nodes": {
			Type:        hcl.TypeSet,
			Description: "A list of synthetic nodes belonging to the location. \n\n You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"availability_location_outage": {
			Type:        hcl.TypeBool,
			Description: "The alerting of location outage is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"availability_node_outage": {
			Type:        hcl.TypeBool,
			Description: "The alerting of node outage is enabled (`true`) or disabled (`false`). \n\n If enabled, the outage of *any* node in the location triggers an alert",
			Optional:    true,
		},
		"location_node_outage_delay_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "Alert if the location or node outage lasts longer than *X* minutes. \n\n Only applicable when **availability_location_outage** or **availability_node_outage** is set to `true`",
			Optional:    true,
		},
		"availability_notifications_enabled": {
			Type:        hcl.TypeBool,
			Description: "The notifications of location and node outage is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"deployment_type": {
			Type:        hcl.TypeString,
			Description: "The deployment type of the location: \n\n* `STANDARD`: The location is deployed on Windows or Linux.\n* `KUBERNETES`: The location is deployed on Kubernetes",
			Optional:    true,
		},
		"auto_update_chromium": {
			Type:        hcl.TypeBool,
			Description: "Auto upgrade of Chromium is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
	}
}

func (me *PrivateSyntheticLocation) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	if me.CountryCode != nil {
		result["country_code"] = me.CountryCode
	}
	if me.RegionCode != nil {
		result["region_code"] = me.RegionCode
	}
	if me.City != nil {
		result["city"] = me.City
	}
	result["latitude"] = me.Latitude
	result["longitude"] = me.Longitude
	if len(me.Nodes) > 0 {
		result["nodes"] = me.Nodes
	}
	result["availability_location_outage"] = me.AvailabilityLocationOutage
	result["availability_node_outage"] = me.AvailabilityNodeOutage
	if me.LocationNodeOutageDelayInMinutes != nil {
		result["location_node_outage_delay_in_minutes"] = *me.LocationNodeOutageDelayInMinutes
	} else {
		result["location_node_outage_delay_in_minutes"] = nil
	}
	result["availability_notifications_enabled"] = me.AvailabilityNotificationsEnabled
	if me.DeploymentType != nil {
		result["deployment_type"] = string(*me.DeploymentType)
	} else {
		result["deployment_type"] = nil
	}
	result["auto_update_chromium"] = me.AutoUpdateChromium

	return result, nil
}

func (me *PrivateSyntheticLocation) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = LocationTypes.Private
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("country_code"); ok {
		me.CountryCode = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("region_code"); ok {
		me.RegionCode = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("city"); ok {
		me.City = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("latitude"); ok {
		me.Latitude = float64(value.(float64))
	}
	if value, ok := decoder.GetOk("longitude"); ok {
		me.Longitude = float64(value.(float64))
	}
	if value, ok := decoder.GetOk("nodes"); ok {
		me.Nodes = []string{}
		for _, elem := range value.(hcl.Set).List() {
			me.Nodes = append(me.Nodes, elem.(string))
		}
	}
	if value, ok := decoder.GetOk("availability_location_outage"); ok {
		me.AvailabilityLocationOutage = value.(bool)
	}
	if value, ok := decoder.GetOk("availability_node_outage"); ok {
		me.AvailabilityNodeOutage = value.(bool)
	}
	if value, ok := decoder.GetOk("location_node_outage_delay_in_minutes"); ok {
		me.LocationNodeOutageDelayInMinutes = opt.NewInt(value.(int))
	}
	if value, ok := decoder.GetOk("availability_notifications_enabled"); ok {
		me.AvailabilityNotificationsEnabled = value.(bool)
	}
	if value, ok := decoder.GetOk("deployment_type"); ok {
		me.DeploymentType = DeploymentType(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("auto_update_chromium"); ok {
		me.AutoUpdateChromium = value.(bool)
	}
	return nil
}