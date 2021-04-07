package NormalRSS

import "github.com/go-resty/resty/v2"

type RSSHelper struct {
	client *resty.Client
}

// 对于一般需要代理的 RSS 使用这个
func NormalRSS(httpProxy string) *RSSHelper {
	nRSS := RSSHelper{}
	// Create a Resty Client
	nRSS.client = resty.New()
	// Setting a Proxy Url and Port
	nRSS.client.SetProxy(httpProxy)

	return &nRSS
}

func (r RSSHelper) GetRSSContent(url string) (string, error) {
	resp, err := r.client.R().Get(url)
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

