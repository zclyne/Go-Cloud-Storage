package meta

import "sort"

// FileMeta: 文件元数据结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: 新增/更新文件元数据
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta: 通过sha1获取文件的元数据对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetLastFileMetas: 返回最新上传的count个文件的元数据
func GetLastFileMetas(count int) []FileMeta {
	// 将fileMetas中的所有元数据转换到数组fMetaArray中
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}
	// 将数组fMetaArray按照上传时间排序
	sort.Sort(ByUploadTime(fMetaArray))

	return fMetaArray[0:count]
}

// RemoveFileMeta: 删除文件元数据
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
