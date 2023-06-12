package types

type WebsiteRequest struct {
	Websites []string `json:"websites"`
}

type WebsiteStatus struct {
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type WebsiteStatusMap struct {
	Status map[string]bool
}
