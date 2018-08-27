package dto

// Service ...
type Service struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	Retries        int `json:"retries,omitempty"`
	ConnectTimeout int `json:"connect_timeout,omitempty"`
	WriteTimeout   int `json:"write_timeout,omitempty"`
	ReadTimeout    int `json:"read_timeout,omitempty"`

	Protocol string `json:"protocol,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Path     string `json:"path,omitempty"`

	CreatedAt int `json:"created_at,omitempty"`
	UpdatedAt int `json:"updated_at,omitempty"`
}
