package dto

// NodeInfo ...
type NodeInfo struct {
	NodeID     string `json:"node_id,omitempty"`
	HostName   string `json:"hostname,omitempty"`
	Version    string `json:"version,omitempty"`
	LuaVersion string `json:"lua_version,omitempty"`
	TagLine    string `json:"tagline,omitempty"`

	Plugins struct {
		EnabledInCluster  []string        `json:"enabled_in_cluster,omitempty"`
		AvailableOnServer map[string]bool `json:"available_on_server,omitempty"`
	} `json:"plugins"`

	Timers struct {
		Pending int `json:"pending,omitempty"`
		Running int `json:"running,omitempty"`
	} `json:"timers,omitempty"`

	PRNGSeeds map[string]interface{} `json:"prng_seeds,omitempty"`
}
