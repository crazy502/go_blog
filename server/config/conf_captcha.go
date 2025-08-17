package config

//配置包

type Captcha struct {
	Height   int     `json:"height" yaml:"height"`       //PNG图片高度（px）
	Width    int     `json:"width" yaml:"width"`         //验证码图片宽度（px）
	Length   int     `json:"length" yaml:"length"`       //验证码结果中默认的数字个数
	MaxSkew  float64 `json:"max_skew" yaml:"max_skew"`   //单个数字的最大偏斜因子（绝对值
	DotCount int     `json:"dot_count" yaml:"dot_count"` //背景圆点数量
}
