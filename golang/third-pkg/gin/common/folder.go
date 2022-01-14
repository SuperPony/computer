package common

import (
	"os"
)

// 判断目录是否存在，如果不存在创建目录
//	folder string: 文件目录
func IsExistFolder(folder string) {
	_, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(folder, 0666)
		}
	}

}
