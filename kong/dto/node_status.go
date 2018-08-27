package dto

// NodeStatus ...
type NodeStatus struct {
	Database struct {
		Reachable bool `json:"reachable"`
	} `json:"database"`

	Server struct {
		ConnectionsWritting uint64 `json:"connections_writting"`
		TotalRequests       uint64 `json:"total_requests"`
		ConnectionsHandled  uint64 `json:"connections_handled"`
		ConnectionsAccepted uint64 `json:"connections_accepted"`
		ConnectionsReading  uint64 `json:"connections_reading"`
		ConnectionsActive   uint64 `json:"connections_active"`
		ConnectionsWaiting  uint64 `json:"connections_waiting"`
	} `json:"server"`
}
