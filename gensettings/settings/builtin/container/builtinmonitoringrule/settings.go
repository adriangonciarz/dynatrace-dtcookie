package builtinmonitoringrule

import "github.com/dtcookie/hcl"

type Settings struct {
	IgnoreDockerPauseContainer     bool `json:"ignoreDockerPauseContainer"`     // Disable monitoring of platform internal pause containers in Kubernetes and OpenShift.
	IgnoreKubernetesPauseContainer bool `json:"ignoreKubernetesPauseContainer"` // Disable monitoring of platform internal pause containers in Kubernetes and OpenShift.
	IgnoreOpenShiftBuildPodName    bool `json:"ignoreOpenShiftBuildPodName"`    // Disable monitoring of intermediate containers created during image build.
	IgnoreOpenShiftSdnNamespace    bool `json:"ignoreOpenShiftSdnNamespace"`    // Disable monitoring of platform internal containers in the openshift-sdn namespace.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ignore_docker_pause_container": {
			Type:        hcl.TypeBool,
			Description: "Disable monitoring of platform internal pause containers in Kubernetes and OpenShift.",
			Optional:    true,
		},
		"ignore_kubernetes_pause_container": {
			Type:        hcl.TypeBool,
			Description: "Disable monitoring of platform internal pause containers in Kubernetes and OpenShift.",
			Optional:    true,
		},
		"ignore_open_shift_build_pod_name": {
			Type:        hcl.TypeBool,
			Description: "Disable monitoring of intermediate containers created during image build.",
			Optional:    true,
		},
		"ignore_open_shift_sdn_namespace": {
			Type:        hcl.TypeBool,
			Description: "Disable monitoring of platform internal containers in the openshift-sdn namespace.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"ignore_docker_pause_container":     me.IgnoreDockerPauseContainer,
		"ignore_kubernetes_pause_container": me.IgnoreKubernetesPauseContainer,
		"ignore_open_shift_build_pod_name":  me.IgnoreOpenShiftBuildPodName,
		"ignore_open_shift_sdn_namespace":   me.IgnoreOpenShiftSdnNamespace,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ignore_docker_pause_container":     &me.IgnoreDockerPauseContainer,
		"ignore_kubernetes_pause_container": &me.IgnoreKubernetesPauseContainer,
		"ignore_open_shift_build_pod_name":  &me.IgnoreOpenShiftBuildPodName,
		"ignore_open_shift_sdn_namespace":   &me.IgnoreOpenShiftSdnNamespace,
	})
}
