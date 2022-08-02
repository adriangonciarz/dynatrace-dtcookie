package process

import (
	tagapi "github.com/dtcookie/dynatrace/api/config/topology/tag"
)

// Processes is a list of short representations of process
type Processes []Process

// Process is a short representation of a process
type Process struct {
	EntityId    string       `json:"entityId"`    // The entity ID of the process
	DisplayName string       `json:"displayName"` // The name of the process as displayed in the UI
	Tags        []tagapi.Tag `json:"tags"`        // The list of entity tags
}
