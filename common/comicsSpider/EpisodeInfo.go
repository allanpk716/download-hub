package comicsSpider

type ComicInfo struct {
	Name		string
	Url			string
	Score		float32
	MaxScore	float32
	TotalVotes	int
	Status		int			// 1，完结。2，更新中
	Classifies	[]string	// 分类关键词
	Eps			[]EpisodeInfo
}

// 漫画中每一话的信息
type EpisodeInfo struct {
	EpName   string     // 第几话，首次使用只需要填写
	MaxPages int        // 一共多少页
	Pages    []PageInfo // 每一页的信息
	Url      string		// 首次使用只需要填写
}
// 漫画一话中一页的信息
type PageInfo struct {
	EpName   	string     // 第几话
	Index 		int			// 第几页
	URL			string		// 这一页的 URL 地址
}