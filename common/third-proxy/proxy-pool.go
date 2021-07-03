package third_proxy

type ProxyPoolProxyInfo struct {
	CheckCount int    `json:"check_count"`
	FailCount  int    `json:"fail_count"`
	LastStatus int    `json:"last_status"`
	LastTime   string `json:"last_time"`
	Proxy      string `json:"proxy"`
	Region     string `json:"region"`
	Source     string `json:"source"`
	Type       string `json:"type"`
}
