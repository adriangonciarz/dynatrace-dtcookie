package processgroup

import (
	tagapi "github.com/dtcookie/dynatrace/api/config/topology/tag"
)

// ProcessGroups is a list of short representations of process group
type ProcessGroups []ProcessGroup

// ProcessGroup is a short representation of a process group
type ProcessGroup struct {
	EntityId    string       `json:"entityId"`    // The entity ID of the process group
	DisplayName string       `json:"displayName"` // The name of the process group as displayed in the UI
	Tags        []tagapi.Tag `json:"tags"`        // The list of entity tags
}
