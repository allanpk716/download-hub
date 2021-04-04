package ho5hoSpider

import (
	"errors"
	"github.com/allanpk716/Downloadhub/common"
	"github.com/go-rod/rod/lib/utils"
	"log"
	"time"
)

type ho5hoSpider struct {
	timeOut time.Duration
	maxRetryTimes int
}

func Newho5hoSpider(timeOut time.Duration, maxRetryTimes int) *ho5hoSpider {
	ho := ho5hoSpider{
		timeOut: timeOut,
		maxRetryTimes: maxRetryTimes,
	}
	return &ho
}

// 一个动漫的所有集的地址
func (h ho5hoSpider) GetAllEpisode() {

}

// 一集里面所有的页地址
func (h ho5hoSpider) GetOneEpisodePicURLs(desURL, httpProxyURL string) error {
	page, err := common.LoadPage(desURL, httpProxyURL, h.timeOut, h.maxRetryTimes)
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
func (h ho5hoSpider) GetOnePic(desURL, httpProxyURL string) error {

	page, err := common.LoadPage(desURL, httpProxyURL, h.timeOut, h.maxRetryTimes)
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