package config

// Email 邮箱配置，可以登录 QQ 邮箱，打开设置，开启第三方服务服务，详情请见 https://mail.qq.com/

type Email struct {
	Host     string `json:"host" yaml:"host"`         // 邮箱服务器地址 例如：smtp.qq.com
	Port     int    `json:"port" yaml:"port"`         // 邮箱服务器端口 例如：465
	From     string `json:"from" yaml:"from"`         // 发件人邮箱地址
	Nickname string `json:"nickname" yaml:"nickname"` // 发件人昵称  用于显示发件人名称
	Secret   string `json:"secret" yaml:"secret"`     //发件人邮箱的密码或应用专用密码，用于身份验证
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`     // 是否使用 SSL 加密连接
}

/*
核心功能：需要一个邮件发送功能来发送验证码、通知等邮件。
支持各种主流邮件服务商，只需正确配置对应参数即可。
*/
