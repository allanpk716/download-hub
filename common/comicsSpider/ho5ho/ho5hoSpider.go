package ho5hoSpider

import (
	"errors"
	"github.com/allanpk716/Downloadhub/common"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
	"log"
	"time"
)

type Ho5hoSpider struct {
	timeOut time.Duration
	maxRetryTimes int
}

func NewHo5hoSpider(timeOut time.Duration, maxRetryTimes int) *Ho5hoSpider {
	ho := Ho5hoSpider{
		timeOut: timeOut,
		maxRetryTimes: maxRetryTimes,
	}
	return &ho
}

// 一个动漫的所有集的地址
func (h Ho5hoSpider) GetAllEpisode() {

}

// 一集里面所有的页地址
func (h Ho5hoSpider) GetOneEpisodePicURLs(urlExInfo common.UrlExInfo) error {
	var err error
	var page *rod.Page
	if urlExInfo.RemoteDockerURL == "" {
		page, err = common.LoadPage(urlExInfo.URL, urlExInfo.HttpProxyURL, h.timeOut, h.maxRetryTimes)

	} else {
		page, err = common.LoadPageFromRemoteDocker(urlExInfo.URL, urlExInfo.HttpProxyURL, urlExInfo.RemoteDockerURL, h.timeOut, h.maxRetryTimes)
	}
	if err != nil {
		return err
	}
	selection, err := page.Element("#single-pager")
	if err != nil {
		return err
	}
	opts, err := selection.Elements("option")
	if err != nil {
		return err
	}
	for _, opt := range opts {
		log.Printf(
			"opt %s : '%s'",
			opt.MustText(),
			*opt.MustAttribute("data-redirect"),
		)
	}

	return nil
}

// 下载一页
func (h Ho5hoSpider) GetOnePic(urlExInfo common.UrlExInfo) error {

	var err error
	var page *rod.Page
	if urlExInfo.RemoteDockerURL == "" {
		page, err = common.LoadPage(urlExInfo.URL, urlExInfo.HttpProxyURL, h.timeOut, h.maxRetryTimes)

	} else {
		page, err = common.LoadPageFromRemoteDocker(urlExInfo.URL, urlExInfo.HttpProxyURL, urlExInfo.RemoteDockerURL, h.timeOut, h.maxRetryTimes)
	}
	if err != nil {
		return err
	}
	el, err := page.Element("#image-0")
	if err != nil {
		return err
	}
	reBytes, err := el.Resource()
	if err != nil {
		return err
	}
	if len(reBytes) < PicMinBytes {
		return ErrDownloadPicIsTooSmall
	}
	err = utils.OutputFile("b.png", reBytes)
	if err != nil {
		return err
	}
	return nil
}

var (
	ErrDownloadPicIsTooSmall = errors.New("download pic is too small")
)
const PicMinBytes = 1024