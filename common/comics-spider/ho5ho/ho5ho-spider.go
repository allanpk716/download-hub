package ho5hoSpider

import (
	"errors"
	"fmt"
	"github.com/allanpk716/Downloadhub/common"
	"github.com/allanpk716/Downloadhub/common/comics-spider"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Ho5hoSpider struct {
	ComicName       string
	saveRootPath    string
	timeOut         time.Duration
	maxRetryTimes   int
	HttpProxyURL    string
	RemoteDockerURL string
	browser			*rod.Browser
}

func NewHo5hoSpider(saveRootPath string,
	HttpProxyURL, RemoteDockerURL string,
	timeOut time.Duration, maxRetryTimes int) (*Ho5hoSpider, error) {
	ho := Ho5hoSpider{
		saveRootPath:  saveRootPath,
		HttpProxyURL: HttpProxyURL,
		RemoteDockerURL: RemoteDockerURL,
		timeOut:       timeOut,
		maxRetryTimes: maxRetryTimes,
	}
	var err error
	if ho.RemoteDockerURL == "" {
		ho.browser, err = common.NewBrowser(ho.HttpProxyURL)

	} else {
		ho.browser, err = common.NewBrowserFromDocker(ho.HttpProxyURL, ho.RemoteDockerURL)
	}
	if err != nil {
		return nil, err
	}
	return &ho, nil
}

func (h *Ho5hoSpider) GetAllComicMatchWhatYouWanted() ([]comics_spider.ComicInfo, error) {
	var err error
	var page *rod.Page
	var comics []comics_spider.ComicInfo
	page, err = common.NewPage(h.browser)
	if err != nil {
		return nil, err
	}
	page, err = common.PageNavigate(page, ComicSiteMainUrl, h.timeOut, h.maxRetryTimes)
	if err != nil {
		return nil, err
	}
	// 进入主要的一级
	allInfoEl, err := page.Element(comicMainALlString)
	if err != nil {
		return nil, err
	}
	// 一共多少页
	comicMainAllPages, err := allInfoEl.Element(comicMainAllPagesString)
	if err != nil {
		return nil, err
	}
	tmpComicPage, err := comicMainAllPages.Text()
	if err != nil {
		return nil, err
	}
	result := regNumber.FindAllString(tmpComicPage, -1)
	if len(result) != 2 {
		return nil, ErrAnalysePagesLenNot2
	}
	nowPage, err := strconv.Atoi(result[0])
	if err != nil {
		return nil, err
	}
	println("NowIndex:", nowPage)
	nowMaxPage, err := strconv.Atoi(result[1])
	if err != nil {
		return nil, err
	}
	// 先分析第一页
	oneComicInfo, err := h.getComicList(page)
	if err != nil {
		return nil, err
	}
	comics = append(comics, oneComicInfo...)
	err = page.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println("nowPage comics:", len(oneComicInfo))
	// 开始遍历
	for i := 2; i <= nowMaxPage; i++ {
		println("NowIndex:", i)
		nowIndexString := strconv.Itoa(i)
		nextUrl := ComicSiteSubPage0 + nowIndexString+ ComicSiteSubPage1
		pageTmp, err := common.NewPageNavigate(h.browser, nextUrl, h.timeOut, h.maxRetryTimes)
		if err != nil {
			return nil, err
		}
		oneComicInfo, err := h.getComicList(pageTmp)
		if err != nil {
			return nil, err
		}
		fmt.Println("nowPage comics:", len(oneComicInfo))
		comics = append(comics, oneComicInfo...)
		err = pageTmp.Close()
		if err != nil {
			return nil, err
		}
	}
	return comics, nil
}

func (h Ho5hoSpider) getComicList(page *rod.Page) ([]comics_spider.ComicInfo, error) {
	var comics []comics_spider.ComicInfo
	// 进入主要的一级
	_, err := page.Element(comicMainALlString)
	if err != nil {
		return nil, err
	}
	// 这一页所有的漫画
	allInfoEl, err := page.Elements("div.page-item-detail.manga")
	if err != nil {
		return nil, err
	}

	for _, subElement := range allInfoEl {
		oneAlink, err := subElement.Element("div > a")
		if err != nil {
			return nil, err
		}
		oneComicUrl, err := oneAlink.Attribute("href")
		if err != nil {
			return nil, err
		}
		oneComicTitle, err := oneAlink.Attribute("title")
		if err != nil {
			return nil, err
		}
		comics = append(comics, comics_spider.ComicInfo{
			Name: *oneComicTitle,
			Url: *oneComicUrl,
		})
	}


	return comics, nil
}

// 一个动漫的所有集的地址
func (h *Ho5hoSpider) GetAllEpisode(rootURL string) (*comics_spider.ComicInfo, error) {
	comicInfo := comics_spider.ComicInfo{}
	comicInfo.Eps = []comics_spider.EpisodeInfo{}
	var err error
	var page *rod.Page
	// 每一次重新开一个
	page, err = common.NewPage(h.browser)
	if err != nil {
		return nil, err
	}
	defer page.Close()
	page, err = common.PageNavigate(page, rootURL, h.timeOut, h.maxRetryTimes)
	if err != nil {
		return nil, err
	}
	// 漫画名称
	comicNameEl, err := page.Element(comicNameElString)
	if err != nil {
		return nil, err
	}
	comicInfo.Name, err = comicNameEl.Text()
	if err != nil {
		return nil, err
	}
	h.ComicName = comicInfo.Name
	// 分数，总分，投票数
	scoreThingsEl, err := page.Element(scoreThingsElString)
	if err != nil {
		return nil, err
	}
	scoreThings, err := scoreThingsEl.Text()
	if err != nil {
		return nil, err
	}
	result := regNumber.FindAllString(scoreThings, -1)
	if len(result) != 3 {
		return nil, ErrAnalyseScoreLenNot3
	}
	// 分数
	nowScore, err := strconv.ParseFloat(result[0], 32)
	if err != nil {
		return nil, err
	}
	comicInfo.Score = float32(nowScore)
	// 总数
	nowMaxScore, err := strconv.ParseFloat(result[1], 32)
	if err != nil {
		return nil, err
	}
	comicInfo.MaxScore = float32(nowMaxScore)
	// 投票数
	comicInfo.TotalVotes, err = strconv.Atoi(result[2])
	if err != nil {
		return nil, err
	}
	// 状态，完结、更新中
	statusEl, err := page.Element(statusElString)
	if err != nil {
		return nil, err
	}
	statusString, err := statusEl.Text()
	if err != nil {
		return nil, err
	}
	if strings.Contains(statusString, StatusStringCompleted) == true {
		comicInfo.Status = 1
	} else if strings.Contains(statusString, StatusStringOnGoing) == true {
		comicInfo.Status = 2
	} else {
		return nil, ErrComicStatusReadError
	}
	// 分类关键词
	classifiesEl, err := page.Element(classifiesElString)
	if err != nil {
		return nil, err
	}
	aLinks, err := classifiesEl.Elements("a")
	if err != nil {
		return nil, err
	}
	comicInfo.Classifies = []string{}
	for _, alink := range aLinks {
		nowText, err := alink.Text()
		if err != nil {
			return nil, err
		}
		comicInfo.Classifies = append(comicInfo.Classifies, nowText)
	}
	// 一共有几话
	epsEl, err := page.Element(epsElString)
	if err != nil {
		return nil, err
	}
	lis, err := epsEl.Elements("li")
	if err != nil {
		return nil, err
	}
	epsMap := make(map[string]*EpsServers)
	for _, li := range lis {
		nowA, err := li.Element("a")
		if err != nil {
			return nil, err
		}
		epsNameAndServer, err := nowA.Text()
		if err != nil {
			return nil, err
		}
		epsUrl, err := nowA.Attribute("href")
		if err != nil {
			return nil, err
		}
		tmpStr := strings.SplitN(epsNameAndServer, "-", 2)
		if len(tmpStr) != 2 {
			return nil, ErrEpsNameAndServerSplitLenNot2
		}
		nowEpsName := strings.TrimSpace(tmpStr[1])
		//nowEpsServer := strings.TrimSpace(tmpStr[0])
		_, ok := epsMap[nowEpsName]
		if ok == false {
			epsMap[nowEpsName] = &EpsServers{
				Urls: []string{},
			}
		}
		epsMap[nowEpsName].Urls = append(epsMap[nowEpsName].Urls, *epsUrl)
	}

	// 这里为了排除 同一集，多个 Server 的观看地址，可能有一个不可用
	// 那么就需要现在进去下载一集的一话
	for k, v := range epsMap {
		for _, url := range v.Urls {
			err = h.GetOneEpisodePicURLs(&comics_spider.EpisodeInfo{
				Url: url,
			})
			if err !=  nil {
				continue
			} else {
				comicInfo.Eps = append(comicInfo.Eps, comics_spider.EpisodeInfo{
					EpName: k,
					Url: url,
				})
			}
		}
	}

	return &comicInfo, nil
}

type EpsServers struct {
	Urls []string
}

// 一集里面所有的页地址
func (h *Ho5hoSpider) GetOneEpisodePicURLs(episodeInfo *comics_spider.EpisodeInfo) error {
	var err error
	var page *rod.Page
	// 每一次重新开一个
	page, err = common.NewPage(h.browser)
	if err != nil {
		return err
	}
	defer page.Close()
	page, err = common.PageNavigate(page, episodeInfo.Url, h.timeOut, h.maxRetryTimes)
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
	episodeInfo.MaxPages = len(opts)
	episodeInfo.Pages = []comics_spider.PageInfo{}
	for index, opt := range opts {
		//picIndex, err := opt.Text()
		//if err != nil {
		//	return err
		//}
		nowUrl, err := opt.Attribute("data-redirect")
		if err != nil {
			return err
		}
		onePage := comics_spider.PageInfo{
			EpName: episodeInfo.EpName,
			Index: index + 1,
			URL: *nowUrl,
		}
		episodeInfo.Pages = append(episodeInfo.Pages, onePage)
	}
	// 尝试读取这一页的图片是否是“裂”的
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

	return nil
}

// 下载一页
func (h *Ho5hoSpider) GetOnePic(pageInfo comics_spider.PageInfo, ExistThenPass bool) error {

	var err error
	// 下载的目标目录是否存在
	desTmpDownloadRootPath := path.Join(h.saveRootPath, h.ComicName, pageInfo.EpName)
	err = os.MkdirAll(desTmpDownloadRootPath, os.ModePerm)
	if err != nil{
		return err
	}
	// 下载目标文件全路径
	desPicFullPath := path.Join(desTmpDownloadRootPath, strconv.Itoa(pageInfo.Index) + common.ComicPicExtetion)
	if ExistThenPass == true {
		if common.Exists(desPicFullPath) == true {
			return nil
		}
	}

	var page *rod.Page
	// 每一次重新开一个
	page, err = common.NewPage(h.browser)
	if err != nil {
		return err
	}
	defer page.Close()
	page, err = common.PageNavigate(page, pageInfo.URL, h.timeOut, h.maxRetryTimes)
	if err != nil {
		return err
	}
	nowImageIndex := strconv.Itoa(pageInfo.Index - 1)
	el, err := page.Element("#image-" + nowImageIndex)
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
	err = utils.OutputFile(desPicFullPath, reBytes)
	if err != nil {
		return err
	}
	return nil
}

func (h *Ho5hoSpider) Close() error {
	return h.browser.Close()
}

var (
	ErrAnalysePagesLenNot2 = errors.New("analyse pages len not 2")
	ErrDownloadPicIsTooSmall = errors.New("download pic is too small")
	ErrAnalyseScoreLenNot3 = errors.New("analyse score len not 3")
	ErrComicStatusReadError = errors.New("comic status read error")
	ErrEpsNameAndServerSplitLenNot2 = errors.New("eps name and server split len not 2")
	regNumber = regexp.MustCompile(`(\-|\+)?\d+(\.\d+)?`)
)
const PicMinBytes = 1024
const StatusStringCompleted = "Completed"
const StatusStringOnGoing = "OnGoing"

const (
	ComicSiteMainUrl = "https://www.ho5ho.com/?m_orderby=views"
	ComicSiteSubPage0 = "https://www.ho5ho.com/page/"
	ComicSiteSubPage1 = "/?m_orderby=views"
	comicMainALlString = "body > div.wrap > div > div > div.c-page-content.style-1 > div > div > div > div.main-col.col-md-8.col-sm-8 > div.main-col-inner > div > div.c-page__content > div.tab-content-wrap > div"
	comicMainAllPagesString = "div.col-12.col-md-12 > div > span.pages"
	comicNameElString   = "body > div.wrap > div > div > div > div.profile-manga.lazy > div > div > div > div.post-title > h1"
	scoreThingsElString = "body > div.wrap > div > div > div > div.profile-manga.lazy > div > div > div > div.tab-summary > div.summary_content_wrap > div > div.post-content > div:nth-child(3) > div.summary-content.vote-details"
	statusElString = "body > div.wrap > div > div > div > div.profile-manga.lazy > div > div > div > div.tab-summary > div.summary_content_wrap > div > div.post-status > div:nth-child(2) > div.summary-content"
	classifiesElString = "body > div.wrap > div > div > div > div.profile-manga.lazy > div > div > div > div.tab-summary > div.summary_content_wrap > div > div.post-content > div:nth-child(8) > div.summary-content > div"
	epsElString = "body > div.wrap > div > div > div > div.c-page-content.style-1 > div > div > div > div.main-col.col-md-8.col-sm-8 > div > div.c-page > div > div.page-content-listing.single-page > div > ul"
)