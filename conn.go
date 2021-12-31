/**
 * @Author: Assassin
 * @Description: 链接模块
 * @File: conn
 * @Version: 1.0.0
 * @Date: 2021/12/24 14:22
 */

package poseidonOss

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Conn struct {
	Url    string
	client *resty.Client
}

/**
 * @Author Assassin
 * @Date 2021/12/24 16:41
 * @Description 初始化clent
 **/
func ConnInitClient() *resty.Client {
	client := resty.New()
	client.SetAllowGetMethodPayload(true)

	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.StatusCode() != http.StatusOK {
			return fmt.Errorf("request statusCode=%v", response.StatusCode())
		}
		return nil
	})
	return client
}

/**
 * @Author Assassin
 * @Date 2021/12/24 16:41
 * @Description 新建http
 **/
func (this *Conn) NewHttp() (httpRequest *resty.Request) {
	httpRequest = this.client.R()
	return
}
