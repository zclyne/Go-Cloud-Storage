package meta

import (
	mydb "cloud-storage/db"
	"sort"
)

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

// UpdateFileMetaDB: 新增/更新文件元数据到mysql数据库中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta: 通过sha1获取文件的元数据对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetFileMetaDB: 从mysql获取文件元数据
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
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
