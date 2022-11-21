package locations

type LocationCollectionElements []LocationCollectionElement

// LocationCollectionElement represents a synthetic location
type LocationCollectionElement struct {
	Name          string         `json:"name"`                    // The name of the location
	ID            *string        `json:"entityId"`                // The Dynatrace entity ID of the location
	Type          LocationType   `json:"type"`                    // The type of the location
	CloudPlatform *CloudPlatform `json:"cloudPlatform,omitempty"` // The cloud provider where the location is hosted. Only applicable to `PUBLIC` locations
	IPs           []string       `json:"ips,omitempty"`           // The list of IP addresses assigned to the location. Only applicable to `PUBLIC` locations
	Stage         *Stage         `json:"stage,omitempty"`         // The release stage of the location
}
