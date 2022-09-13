package common

type SettingsObjectList struct {
	Items []*SettingsObjectListItem `json:"items"`
}

type SettingsObjectListItem struct {
	ObjectID string `json:"objectId"`
}

type SettingsObjectUpdate struct {
	SchemaVersion string      `json:"schemaVersion"`
	Value         interface{} `json:"value"`
}

type SettingsObjectCreate struct {
	SchemaVersion string      `json:"schemaVersion"`
	SchemaID      string      `json:"schemaId"`
	Scope         string      `json:"scope"`
	Value         interface{} `json:"value"`
}

type SettingsObjectResponse struct {
	ObjectID string `json:"objectId"` // The ID of the settings object
	Code     int32  // The HTTP status code for the object
}

type SettingsObjectErrorResponse struct {
	InvalidValue map[string]interface{} `json:"invalidValue,omitempty"` // The value of the setting. \n\n It defines the actual values of settings' parameters. \n\nThe actual content depends on the object's schema.
	Error        *Error                 `json:"error,omitempty"`        // Error details
	Code         *int32                 `json:"code,omitempty"`         // The HTTP status code for the object
}

type Error struct {
	ConstraintViolations []*ConstraintViolation `json:"constraintViolations,omitempty"` // A list of constraint violations
	Message              string                 `json:"message,omitempty"`              // The error message
	Code                 int32                  `json:"code,omitempty"`                 // The HTTP status code
}

type ConstraintViolation struct {
	ParmeterLocation *ParameterLocation `json:"parameterLocation,omitempty"`
	Location         *string            `json:"location,omitempty"`
	Message          *string            `json:"message,omitempty"`
	Path             *string            `json:"path,omitempty"`
}

type ParameterLocation string

var ParameterLocations = struct {
	Path        ParameterLocation
	PayloadBody ParameterLocation
	Query       ParameterLocation
}{
	Path:        ParameterLocation("PATH"),
	PayloadBody: ParameterLocation("PAYLOAD_BODY"),
	Query:       ParameterLocation("QUERY"),
}
