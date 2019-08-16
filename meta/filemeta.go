package meta

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSh1 string
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

// GetFileMeta:通过sha1值获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}