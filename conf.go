/**
 * @Author: Assassin
 * @Description: 客户端配置模块
 * @File: conf
 * @Version: 1.0.0
 * @Date: 2021/12/24 14:12
 */

package poseidon_oss

type Config struct {
	AppId string                 // AppId
	AccessKey    string          // AccessKey
	AccessSecret string          // AccessSecret
	BucketId     int64           // Bucket
	AuthVersion  AuthVersionType //  v1 or v2 signature,default is v1
}

/**
 * @Author Assassin
 * @Date 2021/12/24 14:17
 * @Description 获取默认oss配置
 **/
func getDefaultOssConfig() *Config {
	config := Config{}

	config.AccessKey = ""
	config.AccessSecret = ""
	config.AuthVersion = AuthV1

	return &config
}
