package ogp

// OpenGraph ... OGPでよく使うもの
type OpenGraph struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	FaviconURL  string `json:"favicon_url"`
}
