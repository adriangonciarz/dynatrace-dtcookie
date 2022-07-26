package service

// Services is a list of short representations of services
type Services []Service

// Service is a short representation of a service
type Service struct {
	EntityId    string `json:"entityId"`    // The entity ID of the service
	DisplayName string `json:"displayName"` // The name of the service as displayed in the UI
	Tags        []Tag  `json:"tags"`        // The list of entity tags
}

// Tag is a single definition of a tag used for a Dynatrace entity
type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
