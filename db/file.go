package db

import (
	mydb "cloud-storage/db/mysql"
	"database/sql"
	"fmt"
)

// OnFileUploadFinished: 将文件元数据插入到mysql数据库中，插入成功则返回true，失败则返回false
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	// 使用prepare来构建sql语句可以防止sql注入攻击
	// status的默认值为1
	stmt, err := mydb.DBConn().Prepare(
		"insert into tbl_file (`file_sha1`, `file_name`, `file_size`, `file_addr`, `status`) values (?, ?, ?, ?, 1)")
	if err != nil {
		fmt.Printf("Failed to prepare statement, err:%s\n", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 { // 没有产生新的表记录
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// GetFileMeta: 从mysql获取文件元数据
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_name, file_size, file_addr from tbl_file where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	// 执行数据库查询，并填充tfile，注意Scan中参数顺序要和mysql查询语句的查询顺序相同
	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileName, &tfile.FileSize, &tfile.FileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}
