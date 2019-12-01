package locations

// Country struct
type Country struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	TimeZone       string `json:"time_zone"`
	GeoInformation GeoInformation
	State          []State
}

// GeoInformation struct
type GeoInformation struct {
	Location GeoLocations `json:"location"`
}

// GeoLocations struct
type GeoLocations struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// State struct
type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
