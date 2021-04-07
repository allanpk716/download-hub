package ho5hoSpider

import (
	"fmt"
	"testing"
	"time"
)



func TestAll(t *testing.T) {
	httpProxyURL := "http://127.0.0.1:10809"
	//remoteDockerURL := "ws://192.168.50.135:9222"
	saveRootPath := "X:\\ho5ho\\"
	// 排除不想看的关键词
	noText := make(map[string]int)
	noText["媽媽"]=0
	noText["崩壞"]=0
	noText["幼齒"]=0
	noText["母親"]=0
	noText["母豬"]=0
	noText["阿黑顏"]=0
	noText["黑肉"]=0
	noText["亂倫"]=0
	noText["母子"]=0
	// 实例化
	ho, err := NewHo5hoSpider(saveRootPath,
		httpProxyURL,
		"",
		30 * time.Second, 5)
	if err != nil {
		t.Fatal(err)
	}

	allComics, err := ho.GetAllComicMatchWhatYouWanted()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("All found:", len(allComics))
	for _, comic := range allComics {
		println("Start", comic.Name)
		nowComicInfo, err :=  ho.GetAllEpisode(comic.Url)
		if err != nil {
			t.Fatal(err)
		}
		needPass := false
		fmt.Println("Classifies", nowComicInfo.Classifies)
		for _, classify := range nowComicInfo.Classifies {
			_, ok :=noText[classify]
			if ok == true {
				needPass = true
				break
			}
		}
		if needPass == true {
			println("Pass", comic.Name)
			continue
		}
		for epIndex, ep := range nowComicInfo.Eps {
			println("Ep:", ep.EpName, epIndex+1, "/", len(nowComicInfo.Eps))
			err = ho.GetOneEpisodePicURLs(&ep)
			if err != nil {
				t.Fatal(err)
			}
			for _, page := range ep.Pages {
				println("Page:", page.Index, "/", ep.MaxPages)
				err = ho.GetOnePic(page, true)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
	}
}