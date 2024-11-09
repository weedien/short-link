package mail

// Config 邮件配置
//type Config struct {
//	SMTPHost  string // SMTP服务器地址
//	SMTPPort  int    // SMTP服务器端口
//	Username  string // 发件人邮箱
//	AuthToken string // 发件人密码或授权码
//}

// EmailMessage 邮件消息
type EmailMessage struct {
	From        string   // 发件人
	To          []string // 收件人列表
	Subject     string   // 邮件主题
	Body        string   // 邮件正文（支持HTML格式）
	Attachments []string // 附件路径列表
}

// EmailSender 邮件发送接口
type EmailSender interface {
	SendEmail(msg EmailMessage) error
}
