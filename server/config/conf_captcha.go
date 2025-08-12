package config

type Captcha struct {
	Secret string `json:"secret"`
	Size   int    `json:"size"`
	Expire int    `json:"expire"`
}
