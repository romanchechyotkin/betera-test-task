package nasa

type Metadata struct {
	Title          string `json:"title"`
	Explanation    string `json:"explanation"`
	URL            string `json:"url"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Date           string `json:"date"`
}
