package host

import (
	tagapi "github.com/dtcookie/dynatrace/api/config/topology/tag"
)

// Hosts is a list of short representations of hosts
type Hosts []Host

// Host is a short representation of a host
type Host struct {
	EntityId    string       `json:"entityId"`    // The entity ID of the host
	DisplayName string       `json:"displayName"` // The name of the host as displayed in the UI
	Tags        []tagapi.Tag `json:"tags"`        // The list of entity tags
}
