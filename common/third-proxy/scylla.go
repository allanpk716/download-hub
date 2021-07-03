package third_proxy

type ScyllaProxyInfo struct {
	Proxies []struct {
		Id            int     `json:"id"`
		Ip            string  `json:"ip"`
		Port          int     `json:"port"`
		IsValid       bool    `json:"is_valid"`
		CreatedAt     int     `json:"created_at"`
		UpdatedAt     int     `json:"updated_at"`
		Latency       float64 `json:"latency"`
		Stability     float64 `json:"stability"`
		IsAnonymous   bool    `json:"is_anonymous"`
		IsHttps       bool    `json:"is_https"`
		Attempts      int     `json:"attempts"`
		HttpsAttempts int     `json:"https_attempts"`
		Location      string  `json:"location"`
		Organization  string  `json:"organization"`
		Region        string  `json:"region"`
		Country       string  `json:"country"`
		City          string  `json:"city"`
	} `json:"proxies"`
	Count     int `json:"count"`
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
}
