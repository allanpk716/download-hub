package model

type(
	ProxyConf   []ProxyConfig

	ProxyConfig struct {
		Name string 						// 代理的名称
		Url string							// 代理的地址和端口
		IsStable bool `json:",default=false,options=true|false"`	// 这个仅针对是自建代理需要设置为 True，且设置的顺序影响选择的优先级
	}
)