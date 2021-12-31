/**
 * @Author: Assassin
 * @Description:
 * @File: main_test
 * @Version: 1.0.0
 * @Date: 2021/12/28 9:56
 */

package poseidon_oss

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @Author Assassin
 * @Date 2021/12/28 10:09
 * @Description 测试sdk
 **/
func TestOssSDKUpload(t *testing.T) {
	file, err := os.Open("example.png")
	if err != nil {
		t.Errorf("TestOssSDKUpload os.Open err=%s", err.Error())
		return
	}
	// 传入分配给应用的ak、as、bucketId
	// 如果还没有 找oss管理员创建一个
	client, _ := New("ak", "as", 1, "http://www.baidu.com/upload")
	// file是文件类型  并带上文件名
	res, err := client.UploadFile(file, "example.png")
	if err != nil {
		t.Errorf("TestOssSDKUpload err=%s", err.Error())
		return
	}
	resAssert := assert.New(t)
	resAssert.Equal(res.ErrNo, int64(0))
	if res.ErrNo == 0 {
		t.Logf("%v", res)
	}
	t.Errorf("TestOssSDKUpload err=%v", res)
}
