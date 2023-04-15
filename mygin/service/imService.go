package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func AddIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		fmt.Println(err.Error())
	}
	result := db.Create(&im)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func UpdateIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	result := db.Model(&im).Where("id=?", im.ID).Save(&im)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetIm(c *gin.Context) {
	var ims []model.Im
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	result := db.Find(&ims)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": ims,
	})
}

func DeleteIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Table("ims").Delete(&im)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}
}

func ChangeLines(c *gin.Context) {
	var lines []model.Line
	imid := c.Param("imid")
	err := c.ShouldBind(&lines)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(lines)
	filename := getfilename(imid)
	statusCode := savetocos(lines, filename)
	c.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"msg":  "",
	})
}

// func savetofile(lines []model.Line, filename string) {
// 	data := &model.Datas{
// 		Data: model.Data{
// 			Lines: lines,
// 		},
// 		Ok: true,
// 	}
// 	bytes, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(string(bytes))
// 	log.Println(filename)
// 	path := "/tmp/cosfs/configs/"
// 	file, err := os.OpenFile(path+filename, os.O_TRUNC|os.O_WRONLY, 0777)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer file.Close()
// 	w := bufio.NewWriter(file)
// 	w.WriteString(string(bytes))
// 	w.Flush()
// }

func savetocos(lines []model.Line, filename string) int {
	data := &model.Datas{
		Data: model.Data{
			Lines: lines,
		},
		Ok: true,
	}
	jsonbytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonbytes))
	log.Println(filename)
	client := myutils.CosClient()
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "application/json",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
		},
	}
	objKey := fmt.Sprintf("configs/%v", filename)
	resp, err1 := client.Object.Put(context.Background(), objKey, bytes.NewReader(jsonbytes), opt)
	if err1 != nil {
		log.Println(err1.Error())
	}
	fmt.Println("Object updated successfully.")
	return resp.StatusCode
}

func getfilename(imid string) string {
	client, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	var im model.Im
	db := client.Where("id=?", imid).Find(&im)
	if db.Error != nil {
		log.Println(db.Error.Error())
	}
	s := strings.Split(im.Jsonconfig, "/")
	filename := s[len(s)-1]
	return filename
}
