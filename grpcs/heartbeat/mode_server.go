package heartbeat

type HeartbeatData struct {
	Router string              `json:"router"`
	Params []map[string]string `json:"heartbeat_data"`
}

type ReceiveHeartbeatData struct {
	Host          string          `json:"host"`
	HeartbeatData []HeartbeatData `json:"heartbeat_data"`
}
