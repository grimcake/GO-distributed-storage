package meta

import (
	mydb "filestore-server/db"
	"sort"
)

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSh1  string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// fileMetas：sha1和FileMeta的映射
var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSh1] = fmeta
}

// UpdateFileMetaDB:新增、更新元信息到mysql中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(
		fmeta.FileSh1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta:通过sha1值获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetLatFileMetas:获取批量的文件元信息列表
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

// RemoveFileMeta:删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
