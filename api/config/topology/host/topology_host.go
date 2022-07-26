package host

// Hosts is a list of short representations of hosts
type Hosts []Host

// Host is a short representation of a host
type Host struct {
	EntityId    string `json:"entityId"`    // The entity ID of the host
	DisplayName string `json:"displayName"` // The name of the host as displayed in the UI
	Tags        []Tag  `json:"tags"`        // The list of entity tags
}

// Tag is a single definition of a tag used for a Dynatrace entity
type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
