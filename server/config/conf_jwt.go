package config

// Jwt jwt 配置
type Jwt struct {
	AccessTokenSecret      string `json:"access_token_secret" yaml:"access_token_secret"`             // 用于生成和验证访问令牌的密钥
	RefreshTokenSecret     string `json:"refresh_token_secret" yaml:"refresh_token_secret"`           // 用于生成和验证刷新令牌的密钥
	AccessTokenExpiryTime  string `json:"access_token_expiry_time" yaml:"access_token_expiry_time"`   // 访问令牌的过期时间，例如 "15m" 表示 15 分钟
	RefreshTokenExpiryTime string `json:"refresh_token_expiry_time" yaml:"refresh_token_expiry_time"` // 刷新令牌的过期时间，例如 "30d" 表示 30 天
	Issuer                 string `json:"issuer" yaml:"issuer"`                                       // JWT 的签发者信息，通常是应用或服务的名称
}

/*
JWT 是一种开放标准（RFC 7519），用于在各方之间安全地传输信息。在博客项目中，JWT 主要用于：
	用户身份验证 - 验证用户是否已登录
	权限控制 - 确定用户可以访问哪些资源
	会话管理 - 维持用户的登录状态
*/
//1. 用户登录 → 服务器验证凭据 → 生成 JWT 令牌
//2. 服务器将令牌返回给客户端（通常在响应头或响应体中）
//3. 客户端存储令牌（通常在 localStorage 或 cookie 中）
//4. 客户端后续请求时在请求头中携带 JWT 令牌
//5. 服务器验证令牌有效性 → 允许或拒绝请求
