package ipmappings

import "github.com/dtcookie/hcl"

type Settings struct {
	City        *string `json:"city"`        // The city name of the location.
	CountryCode string  `json:"countryCode"` // The country code of the location. \n\n Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
	Ip          string  `json:"ip"`          // Single IP or IP range start address
	IpTo        *string `json:"ipTo"`        // IP range end
	Latitude    float64 `json:"latitude"`    // Latitude
	Longitude   float64 `json:"longitude"`   // Longitude
	RegionCode  *string `json:"regionCode"`  // The region code of the location. \n\n For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix. \n\n For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"city": {
			Type:        hcl.TypeString,
			Description: "The city name of the location.",
			Optional:    true,
		},
		"country_code": {
			Type:        hcl.TypeString,
			Description: "The country code of the location. \n\n Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).",
			Required:    true,
		},
		"ip": {
			Type:        hcl.TypeString,
			Description: "Single IP or IP range start address",
			Required:    true,
		},
		"ip_to": {
			Type:        hcl.TypeString,
			Description: "IP range end",
			Optional:    true,
		},
		"latitude": {
			Type:        hcl.TypeFloat,
			Description: "Latitude",
			Required:    true,
		},
		"longitude": {
			Type:        hcl.TypeFloat,
			Description: "Longitude",
			Required:    true,
		},
		"region_code": {
			Type:        hcl.TypeString,
			Description: "The region code of the location. \n\n For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix. \n\n For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"city":         me.City,
		"country_code": me.CountryCode,
		"ip":           me.Ip,
		"ip_to":        me.IpTo,
		"latitude":     me.Latitude,
		"longitude":    me.Longitude,
		"region_code":  me.RegionCode,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"city":         &me.City,
		"country_code": &me.CountryCode,
		"ip":           &me.Ip,
		"ip_to":        &me.IpTo,
		"latitude":     &me.Latitude,
		"longitude":    &me.Longitude,
		"region_code":  &me.RegionCode,
	})
}
