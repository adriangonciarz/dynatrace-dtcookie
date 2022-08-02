package application

import (
	tagapi "github.com/dtcookie/dynatrace/api/config/topology/tag"
)

// Applications is a list of short representations of applications
type Applications []Application

// Application is a short representation of an application
type Application struct {
	EntityId    string       `json:"entityId"`    // The entity ID of the application
	DisplayName string       `json:"displayName"` // The name of the application as displayed in the UI
	Tags        []tagapi.Tag `json:"tags"`        // The list of entity tags
}
