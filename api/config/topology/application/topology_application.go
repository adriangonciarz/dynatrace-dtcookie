package application

// Applications is a list of short representations of applications
type Applications []Application

// Application is a short representation of an application
type Application struct {
	EntityId    string `json:"entityId"`    // The entity ID of the application
	DisplayName string `json:"displayName"` // The name of the application as displayed in the UI
	Tags        []Tag  `json:"tags"`        // The list of entity tags
}

// Tag is a single definition of a tag used for a Dynatrace entity
type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
