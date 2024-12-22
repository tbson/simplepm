package infra

type Config struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}

type Hook struct {
	Type          string   `json:"type"`
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Active        bool     `json:"active"`
	Events        []string `json:"events"`
	Config        Config   `json:"config"`
	UpdatedAt     string   `json:"updated_at"`
	CreatedAt     string   `json:"created_at"`
	AppID         int      `json:"app_id"`
	DeliveriesURL string   `json:"deliveries_url"`
}

type InputData struct {
	Zen    string `json:"zen"`
	HookID int    `json:"hook_id"`
	Hook   Hook   `json:"hook"`
}
