package technology

import "github.com/dtcookie/hcl"

type Settings struct {
	BoshProcessManager bool `json:"boshProcessManager"` // Platform: Cloud Foundry\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.159
	Containerd         bool `json:"containerd"`         // Platform: Kubernetes\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.169
	Crio               bool `json:"crio"`               // Platform: Kubernetes\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.163
	Docker             bool `json:"docker"`             // Platform: Docker and Kubernetes\n\nStatus: Released\n\nOperating system: Linux
	DockerWindows      bool `json:"dockerWindows"`      // Platform: Docker\n\nStatus: Early adopter\n\nOperating system: Windows\n\nMin agent version: 1.149
	Garden             bool `json:"garden"`             // Platform: Cloud Foundry\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.133
	Winc               bool `json:"winc"`               // Platform: Cloud Foundry\n\nStatus: Early adopter\n\nOperating system: Windows\n\nMin agent version: 1.175
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"bosh_process_manager": {
			Type:        hcl.TypeBool,
			Description: "Platform: Cloud Foundry\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.159",
			Optional:    true,
		},
		"containerd": {
			Type:        hcl.TypeBool,
			Description: "Platform: Kubernetes\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.169",
			Optional:    true,
		},
		"crio": {
			Type:        hcl.TypeBool,
			Description: "Platform: Kubernetes\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.163",
			Optional:    true,
		},
		"docker": {
			Type:        hcl.TypeBool,
			Description: "Platform: Docker and Kubernetes\n\nStatus: Released\n\nOperating system: Linux",
			Optional:    true,
		},
		"docker_windows": {
			Type:        hcl.TypeBool,
			Description: "Platform: Docker\n\nStatus: Early adopter\n\nOperating system: Windows\n\nMin agent version: 1.149",
			Optional:    true,
		},
		"garden": {
			Type:        hcl.TypeBool,
			Description: "Platform: Cloud Foundry\n\nStatus: Released\n\nOperating system: Linux\n\nMin agent version: 1.133",
			Optional:    true,
		},
		"winc": {
			Type:        hcl.TypeBool,
			Description: "Platform: Cloud Foundry\n\nStatus: Early adopter\n\nOperating system: Windows\n\nMin agent version: 1.175",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"bosh_process_manager": me.BoshProcessManager,
		"containerd":           me.Containerd,
		"crio":                 me.Crio,
		"docker":               me.Docker,
		"docker_windows":       me.DockerWindows,
		"garden":               me.Garden,
		"winc":                 me.Winc,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"bosh_process_manager": &me.BoshProcessManager,
		"containerd":           &me.Containerd,
		"crio":                 &me.Crio,
		"docker":               &me.Docker,
		"docker_windows":       &me.DockerWindows,
		"garden":               &me.Garden,
		"winc":                 &me.Winc,
	})
}
