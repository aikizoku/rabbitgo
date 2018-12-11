package ogp

// OpenGraph ... OGPでよく使うもの
type OpenGraph struct {
	Type        string            `json:"type"`
	URL         string            `json:"url"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	SiteName    string            `json:"site_name"`
	Images      []*OpenGraphImage `json:"images"`
}

// OpenGraphImage ... OGPの画像
type OpenGraphImage struct {
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Type      string `json:"type"`
	Width     uint64 `json:"width"`
	Height    uint64 `json:"height"`
}
