package common

import "os"

const (
	TmpDownloadFloderName = "TmpDownload"
	ComicPicExtetion = ".jpg"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}