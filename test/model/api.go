package model

// API ...
type API struct {
	Name     string
	Overview *APIOverview
	Request  *APIRequest
	Response *APIResponse
}

// APIOverview ...
type APIOverview struct {
	Type string
	URL  *APIOverviewURL
	URI  string
}

// APIOverviewURL ...
type APIOverviewURL struct {
	Staging    string
	Production string
}

// APIRequest ...
type APIRequest struct {
	Method  string
	URI     string
	Headers string
	Params  string
}

// APIResponse ...
type APIResponse struct {
	StatusCode int
	Body       string
}
