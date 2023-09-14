package utils

import (
	"encoding/csv"
	"fmt"
	"micro_shopping/idl/pb"
	"mime/multipart"
)

func ReadCsv(fileHeader *multipart.FileHeader) (*pb.BulkRequest, error) {
	f, err := fileHeader.Open()
	if err != nil {
		fmt.Println("open file error", err)
		return nil, err
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("close file error", err)
		}
	}(f)

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("read error", err)
		return nil, err
	}

	var bulkCategory *pb.BulkRequest
	bulkCategory = new(pb.BulkRequest)
	result := bulkCategory.GetBulkRequest()
	for _, line := range lines[1:] {
		result = append(result, &pb.CategoryRequest{
			Name: line[0],
			Desc: line[1],
		})
	}
	return bulkCategory, nil
}
