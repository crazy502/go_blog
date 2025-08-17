package config

// Gaode 高德地图配置，详情可见https://lbs.amap.com/
type Gaode struct {
	Enable bool   `json:"enable" yaml:"enable"` // 是否启用
	Key    string `json:"key" yaml:"key"`       // 高德服务的应用密码，用于身份验证和服务访问
}
