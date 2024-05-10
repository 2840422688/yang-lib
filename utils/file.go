package Utils

//文件工具类
import (
	"io/fs"
	"os"
)

// 测试
func MKDir(perm fs.FileMode, dirPath string) error {
	if err := os.Mkdir(dirPath, perm); err != nil {
		return err
	}
	return nil
}
func MKFile(fileName string, dirPath string) error {
	if _, err := os.Create(dirPath + fileName); err != nil {
		return err
	}
	return nil
}
