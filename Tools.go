package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

/*map 转json
 */
func MapToJson(s interface{}) string {
	b, err := json.MarshalIndent(s, "", "")
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return ""
	}
	return string(b)
}

/*
*
公共报错返回接口格式
*/
func DoError(c *gin.Context, msg string) {
	mapData := make(map[string]interface{})
	mapData["code"] = http.StatusOK
	mapData["msg"] = msg
	c.JSON(http.StatusOK, mapData)
}

/*
*
 */
func DoData(c *gin.Context, data interface{}) {
	mapData := make(map[string]interface{})
	mapData["code"] = http.StatusOK
	mapData["msg"] = "成功"
	mapData["data"] = data
	c.JSON(http.StatusOK, mapData)
}

// 获取文件的md5码
func GetFileMd5ByPath(filename string) string {
	// 文件全路径名
	path := fmt.Sprintf("./%s", filename)
	pFile, err := os.Open(path)
	if err != nil {
		fmt.Errorf("打开文件失败，filename=%v, err=%v", filename, err)
		return ""
	}
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil))
}

// 获取文件的md5码
func GetFileMd5ByFile(fileMultipart *multipart.FileHeader) string {
	file, err := fileMultipart.Open()
	if err != nil {

	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("计算 MD5 值时发生错误:", err)
		return ""
	}
	md5Sum := hex.EncodeToString(hash.Sum(nil)) // 将哈希值转换为字符串形式
	fmt.Println("文件的 MD5 值:", md5Sum)           // 打印 MD5 值到控制台（此处仅为示例）
	return md5Sum
}
