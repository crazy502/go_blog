package config

type Config struct {
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Email   Email   `json:"email" yaml:"email"`
	ES      ES      `json:"es" yaml:"es"`
	Gaode   Gaode   `json:"gaode" yaml:"gaode"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	Qiniu   Qiniu   `json:"qiniu" yaml:"qiniu"`
	QQ      QQ      `json:"qq" yaml:"qq"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	System  System  `json:"system" yaml:"system"`
	Upload  Upload  `json:"upload" yaml:"upload"`
	Website Website `json:"website" yaml:"website"`
	Zap     Zap     `json:"zap" yaml:"zap"`
}

//把所有的配置写成一个结构体，方便统一管理和使用
//好处：
// 1. 方便统一管理和使用
// 2. 方便修改和维护
// 3. 方便扩展和添加新的配置项
// 4. 方便代码的复用
