package utils

import(
	"strconv"
)

func Pagination(pageStr string, pageSizeStr string) (int, int) {
	var err error
	var page, pagesize int
	page, err = strconv.Atoi(pageStr)
	if(err != nil || page = 0){
		page = 1
	}
	pageSize, err = strconv.Atoi(pageSizeStr)
	if(err != nil || pageSize = 0){
		pageSizeStr = 10
	}
	offset := (page-1)*pageSizeStr
	return offset, page
}