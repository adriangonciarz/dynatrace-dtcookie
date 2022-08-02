package service

import (
	tagapi "github.com/dtcookie/dynatrace/api/config/topology/tag"
)

// Services is a list of short representations of services
type Services []Service

// Service is a short representation of a service
type Service struct {
	EntityId    string       `json:"entityId"`    // The entity ID of the service
	DisplayName string       `json:"displayName"` // The name of the service as displayed in the UI
	Tags        []tagapi.Tag `json:"tags"`        // The list of entity tags
}
