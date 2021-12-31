/**
 * @Author: 陈磊
 * @Date: 2021/9/18 14:44
 */

package sirius_oss

import (
	"crypto/md5"
	"fmt"
)

/**
 * @Author 陈磊
 * @Date 2021/9/22 11:00
 * @Description md5加密
 **/
func CreateMd5(signParam string) (sign string) {
	data := []byte(signParam)
	has := md5.Sum(data)
	sign = fmt.Sprintf("%x", has)
	return
}
