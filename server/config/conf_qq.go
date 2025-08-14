package config

// QQ 用于配置 QQ 登录相关信息,详情请查看 https://developer.work.weixin.qq.com/document/path/91039
type QQ struct {
	Enabled     bool   `json:"enabled" yaml:"enabled"`           // 是否启用 QQ 登录
	AppID       string `json:"app_id" yaml:"app_id"`             // QQ 登录应用 ID
	AppKey      string `json:"app_key" yaml:"app_key"`           // QQ 登录应用 Key
	RedirectURL string `json:"redirect_url" yaml:"redirect_url"` // QQ 登录回调地址
}

func (qq QQ) QQLoginURL() string {
	return "https://graph.qq.com/oauth2.0/authorize?" +
		"response_type=code&" +
		"client_id=" + qq.AppID + "&" +
		"redirect_uri=" + qq.RedirectURL
}
