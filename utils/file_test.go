package Utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMKDir(t *testing.T) {
	os.Remove("./testdir/testfile.txt")
	os.Remove("./testdir")
	os.Remove("./testdir2")
	// 正常情况：创建目录
	err := MKDir(0755, "./testdir/")
	assert.NoError(t, err)

	// 异常情况：创建已存在的目录
	err = MKDir(0755, "./testdir/")
	assert.Error(t, err)

}

func TestMKFile(t *testing.T) {
	os.Remove("./testdir/testfile.txt")
	os.Remove("./testdir")
	os.Remove("./testdir2")
	// 创建目录
	err := MKDir(0755, "./testdir/")
	assert.NoError(t, err)

	//正常情况：创建文件
	err = MKFile("testfile.txt", "./testdir/")
	assert.NoError(t, err)

	os.Remove("./testdir/testfile.txt")
	os.Remove("./testdir")

	// 异常情况：目录不存在
	err = MKFile("testfile.txt", "./nonexistentdir/")
	assert.Error(t, err)
}
