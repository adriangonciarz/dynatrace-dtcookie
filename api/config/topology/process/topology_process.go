package process

// Processes is a list of short representations of process
type Processes []Process

// Process is a short representation of a process
type Process struct {
	EntityId    string `json:"entityId"`    // The entity ID of the process
	DisplayName string `json:"displayName"` // The name of the process as displayed in the UI
	Tags        []Tag  `json:"tags"`        // The list of entity tags
}

// Tag is a single definition of a tag used for a Dynatrace entity
type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
