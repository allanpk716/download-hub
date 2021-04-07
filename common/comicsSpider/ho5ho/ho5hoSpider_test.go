package ho5hoSpider

import (
	"testing"
	"time"
)



func TestAll(t *testing.T) {
	httpProxyURL := "http://127.0.0.1:10809"
	comicUrl := "https://www.ho5ho.com/%E4%B8%AD%E5%AD%97h%E6%BC%AB/%e7%84%a1%e7%a2%bc%e6%88%90%e4%ba%ba%e6%bc%ab%e7%95%ab-%e8%83%b8%e7%bd%a9%e8%a2%ab%e4%ba%ba%e5%81%b7%e4%ba%86%e5%8f%aa%e5%a5%bd%e7%9c%9f%e7%a9%ba%e5%8e%bb%e9%9d%a2%e8%a9%a6/"
	//remoteDockerURL := "ws://192.168.50.135:9222"
	saveRootPath := "X:\\ho5ho\\"

	ho, err := NewHo5hoSpider(saveRootPath,
		httpProxyURL,
		"",
		30 * time.Second, 5)
	if err != nil {
		t.Fatal(err)
	}
	nowComicInfo, err :=  ho.GetAllEpisode(comicUrl)
	if err != nil {
		t.Fatal(err)
	}
	for _, ep := range nowComicInfo.Eps {
		err = ho.GetOneEpisodePicURLs(&ep)
		if err != nil {
			t.Fatal(err)
		}
		for _, page := range ep.Pages {
			err = ho.GetOnePic(page, true)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

}