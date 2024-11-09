package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"shortlink/internal/base"
)

// SMTPMailer SMTP实现
type SMTPMailer struct {
}

// NewSMTPMailer 创建新的SMTPMailer
func NewSMTPMailer() *SMTPMailer {
	return &SMTPMailer{}
}

// SendEmail 实现发送邮件逻辑
func (s *SMTPMailer) SendEmail(msg EmailMessage) error {

	config := base.GetConfig().Email

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", msg.From)
	m.SetHeader("To", msg.To...)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/html", msg.Body) // 支持HTML格式的邮件正文

	// 添加附件（如果有）
	for _, attachment := range msg.Attachments {
		m.Attach(attachment)
	}

	// 配置SMTP拨号器
	dialer := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.Username, config.AuthToken)

	// 发送邮件
	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
