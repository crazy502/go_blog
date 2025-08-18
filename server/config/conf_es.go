package config

// ES Elasticsearch 配置
type ES struct {
	URL            string `json:"url" yaml:"url"`                           // Elasticsearch 服务的 URL，例如 http://localhost:9200
	Username       string `json:"username" yaml:"username"`                 // Elasticsearch 登录用户名
	Password       string `json:"password" yaml:"password"`                 // Elasticsearch 登录密码
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 是否在控制台打印 Elasticsearch 语句，true 表示打印，false 表示不打印
}

/*
在博客项目中，Elasticsearch 可能用于：
全文搜索 - 搜索文章内容、标题、标签等
日志分析 - 分析用户访问日志、行为数据
标签和分类统计 - 统计热门标签、分类等信息
用户行为分析 - 分析用户阅读偏好等
*/
