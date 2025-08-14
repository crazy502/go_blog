package config

// QiniuConfig 七牛云配置 ，详情参见 https://developer.qiniu.com/kodo/manual/1206/storage-classes
type Qiniu struct {
	Zone          string `json:"zone" yaml:"zone"`                       // 存储区域
	Bucket        string `json:"bucket" yaml:"bucket"`                   // 存储空间名
	ImgPath       string `json:"img_path" yaml:"img_path"`               // 图片存储路径，CDN 加速域名下，图片访问路径前缀
	AccessKey     string `json:"access_key" yaml:"access_key"`           // 七牛云 AccessKey
	SecretKey     string `json:"secret_key" yaml:"secret_key"`           // 七牛云 SecretKey
	UseHTTPS      bool   `json:"use_https" yaml:"use_https"`             // 是否使用 HTTPS 协议
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains"` // 是否使用 CDN 加速域名,上传图片时，是否使用 CDN 域名
}
