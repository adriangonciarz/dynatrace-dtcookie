package service

type Services []Service

// Service has no documentation
type Service struct {
	EntityId    string `json:"entityId"`    // The entity ID of the service
	DisplayName string `json:"displayName"` // The name of the service as displayed in the UI
	Tags        []Tag  `json:"tags"`        // The list of entity tags
}

type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
