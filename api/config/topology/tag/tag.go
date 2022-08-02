package tag

// Tag is a single definition of a tag used for a Dynatrace entity
type Tag struct {
	Context string  `json:"context"`
	Key     string  `json:"key"`
	Value   *string `json:"value,omitempty"`
}
