package extension

import "github.com/dtcookie/hcl"

type Settings struct {
	DiskDeviceMonitoringExtensionActive bool `json:"diskDeviceMonitoringExtensionActive"` // The Disk Analytics feature requires an extension to be added to your environment. The Disk Analytics extension consumes custom metrics and [Davis data units](https://www.dynatrace.com/support/help/shortlink/metric-cost-calculation).\n\nAfter you have installed the Disk Analytics extension, you can enable the Data Collection in host or host-group level settings. If you enable the Data Collection without adding the extension the data is only visible in the data explorer.\n\nFor details, see [Disk Analytics extension documentation](https://dt-url.net/3a03v9v).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"disk_device_monitoring_extension_active": {
			Type:        hcl.TypeBool,
			Description: "The Disk Analytics feature requires an extension to be added to your environment. The Disk Analytics extension consumes custom metrics and [Davis data units](https://www.dynatrace.com/support/help/shortlink/metric-cost-calculation).\n\nAfter you have installed the Disk Analytics extension, you can enable the Data Collection in host or host-group level settings. If you enable the Data Collection without adding the extension the data is only visible in the data explorer.\n\nFor details, see [Disk Analytics extension documentation](https://dt-url.net/3a03v9v).",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"disk_device_monitoring_extension_active": me.DiskDeviceMonitoringExtensionActive,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk_device_monitoring_extension_active": &me.DiskDeviceMonitoringExtensionActive,
	})
}
