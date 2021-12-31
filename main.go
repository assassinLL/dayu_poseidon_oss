/**
 * @Author: Assassin
 * @Description:
 * @File: main.go
 * @Version: 1.0.0
 * @Date: 2021/12/24 14:11
 */

package poseidonOss

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/golang-module/carbon"
)

type (
	// Client OSS client
	Client struct {
		Config *Config // OSS client configuration
		Conn   *Conn   // Send HTTP request
	}

	// ClientOption client option such as UseCname, Timeout, SecurityToken.
	ClientOption func(*Client)
)

/**
 * @Author Assassin
 * @Date 2021/12/24 16:18
 * @Description 上传文件返回值
 **/
type UploadFileResponse struct {
	ErrNo   int64  `json:"err_no"`
	Message string `json:"message"`
	Data    data
}

/**
 * @Author Assassin
 * @Date 2021/12/27 14:48
 * @Description 具体数据
 **/
type data struct {
	Ext  string `json:"ext"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

/**
 * @Author Assassin
 * @Date 2021/12/24 14:21
 * @Description 创建新的oss
 **/
func New(accessKey, accessSecret string, BucketId int64, serviceUrl string, options ...ClientOption) (*Client, error) {
	// Configuration
	config := getDefaultOssConfig()
	config.AccessKey = accessKey
	config.AccessSecret = accessSecret
	config.BucketId = BucketId

	// HTTP connect
	conn := &Conn{Url: serviceUrl}
	conn.client = ConnInitClient()

	// OSS client
	client := &Client{
		Config: config,
		Conn:   conn,
	}

	// Client options parse
	for _, option := range options {
		option(client)
	}

	if config.AuthVersion != AuthV1 {
		return nil, fmt.Errorf("oss-SDK Init client Error, invalid Auth version: %v", config.AuthVersion)
	}

	return client, nil
}

/**
 * @Author Assassin
 * @Date 2021/12/24 14:31
 * @Description 上传文件
 **/
func (client Client) UploadFile(file multipart.File, filename string) (res *UploadFileResponse, err error) {
	// 转换为二进制流
	fileBytes := make([]byte, 0)
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		err = fmt.Errorf("oss-SDK UploadFile io.Copy err=%s", err.Error())
		return
	}
	fileBytes = buf.Bytes()

	// 生成参数
	salt := carbon.Now().Timestamp() // 混淆参数盐 为时间戳
	formData := map[string]string{
		"bucket_id":         strconv.FormatInt(client.Config.BucketId, 10),
		"access_key":        client.Config.AccessKey,
		"signature_version": string(client.Config.AuthVersion),
		"salt":              strconv.FormatInt(salt, 10),
		"signature":         getSignature(client, salt),
	}

	// 请求oss服务器
	resp, err := client.Conn.client.NewRequest().
		SetFileReader("file", filename, bytes.NewReader(fileBytes)).
		SetFormData(formData).
		Post(client.Conn.Url)
	if err != nil {
		err = fmt.Errorf("oss-SDK UploadFile request post formData")
		return
	}

	// 解码
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		err = fmt.Errorf("oss-SDK UploadFile json.Unmarshal respBody=%s, err=%s", resp.Body(), err.Error())
		return
	}
	return
}

/**
 * @Author Assassin
 * @Date 2021/12/27 15:00
 * @Description 签名生成
 **/
func getSignature(client Client, salt int64) (signature string) {
	switch client.Config.AuthVersion {
	case AuthV1:
		oriString := fmt.Sprintf("%s%s%d", client.Config.AccessKey, client.Config.AccessSecret, salt)
		signature = CreateMd5(oriString)
	default:
		return
	}
	return
}
