package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	mail "github.com/wneessen/go-mail"
)

const (
	// Password reset token expiry time (15 minutes)
	PasswordResetTokenExpiry = 15 * time.Minute
	// Redis key prefix for password reset tokens
	PasswordResetKeyPrefix = "password_reset:"
)

type EmailServiceInterface interface {
	SendPasswordResetEmail(ctx context.Context, email, token string) error
	GeneratePasswordResetToken(ctx context.Context, userID int64, email string) (string, error)
	ValidatePasswordResetToken(ctx context.Context, token string) (int64, string, error)
	InvalidatePasswordResetToken(ctx context.Context, token string) error
}

type emailService struct {
	redisClient *redis.Client
	smtpHost    string
	smtpPort    int
	smtpUser    string
	smtpPass    string
	fromEmail   string
	fromName    string
	frontendURL string
	logger      *log.Logger
}

type EmailConfig struct {
	RedisClient *redis.Client
	SMTPHost    string
	SMTPPort    int
	SMTPUser    string
	SMTPPass    string
	FromEmail   string
	FromName    string
	FrontendURL string
	Logger      *log.Logger
}

func NewEmailService(cfg EmailConfig) *emailService {
	return &emailService{
		redisClient: cfg.RedisClient,
		smtpHost:    cfg.SMTPHost,
		smtpPort:    cfg.SMTPPort,
		smtpUser:    cfg.SMTPUser,
		smtpPass:    cfg.SMTPPass,
		fromEmail:   cfg.FromEmail,
		fromName:    cfg.FromName,
		frontendURL: cfg.FrontendURL,
		logger:      cfg.Logger,
	}
}

// GeneratePasswordResetToken creates a secure token and stores it in Redis
func (s *emailService) GeneratePasswordResetToken(ctx context.Context, userID int64, email string) (string, error) {
	// Generate a secure random token
	tokenBytes := make([]byte, 3)
	if _, err := rand.Read(tokenBytes); err != nil {
		s.logger.Println("error generating random token:", err)
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	// Store token in Redis with user ID and email
	key := PasswordResetKeyPrefix + token
	data := fmt.Sprintf("%d:%s", userID, email)

	err := s.redisClient.Set(ctx, key, data, PasswordResetTokenExpiry).Err()
	if err != nil {
		s.logger.Println("error storing token in redis:", err)
		return "", err
	}

	s.logger.Printf("generated password reset token for user %d", userID)
	return token, nil
}

// ValidatePasswordResetToken checks if the token is valid and returns the user ID and email
func (s *emailService) ValidatePasswordResetToken(ctx context.Context, token string) (int64, string, error) {
	key := PasswordResetKeyPrefix + token

	data, err := s.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		s.logger.Println("password reset token not found or expired")
		return 0, "", fmt.Errorf("invalid or expired token")
	}
	if err != nil {
		s.logger.Println("error getting token from redis:", err)
		return 0, "", err
	}

	// Parse the data (format: "userID:email")
	var userID int64
	var email string
	_, err = fmt.Sscanf(data, "%d:%s", &userID, &email)
	if err != nil {
		s.logger.Println("error parsing token data:", err)
		return 0, "", err
	}

	return userID, email, nil
}

// InvalidatePasswordResetToken removes the token from Redis
func (s *emailService) InvalidatePasswordResetToken(ctx context.Context, token string) error {
	key := PasswordResetKeyPrefix + token
	err := s.redisClient.Del(ctx, key).Err()
	if err != nil {
		s.logger.Println("error deleting token from redis:", err)
		return err
	}
	s.logger.Println("invalidated password reset token")
	return nil
}

// SendPasswordResetEmail sends an email with the password reset link
func (s *emailService) SendPasswordResetEmail(ctx context.Context, email, token string) error {
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", s.frontendURL, token)

	// Create the email message
	m := mail.NewMsg()
	if err := m.From(s.fromEmail); err != nil {
		s.logger.Println("error setting from address:", err)
		return err
	}
	if err := m.To(email); err != nil {
		s.logger.Println("error setting to address:", err)
		return err
	}

	m.Subject("Reset Your AceThatPaper Password")

	// HTML email body
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: 'Arial Black', Helvetica, sans-serif; background-color: #f5f5f5;">
    <table role="presentation" style="width: 100%%; max-width: 600px; margin: 0 auto; background-color: #ffffff; border: 4px solid #000000;">
        <tr>
            <td style="padding: 40px 30px; text-align: center; background-color: #FFEB3B; border-bottom: 4px solid #000000;">
                <h1 style="margin: 0; font-size: 32px; color: #000000; text-transform: uppercase;">AceThatPaper</h1>
            </td>
        </tr>
        <tr>
            <td style="padding: 40px 30px;">
                <h2 style="margin: 0 0 20px; font-size: 24px; color: #000000;">Password Reset Request</h2>
                <p style="margin: 0 0 20px; font-size: 16px; line-height: 1.6; color: #333333;">
                    We received a request to reset your password. Click the button below to create a new password.
                </p>
                <table role="presentation" style="width: 100%%;">
                    <tr>
                        <td style="text-align: center; padding: 20px 0;">
                            <a href="%s" style="display: inline-block; padding: 16px 40px; font-size: 18px; font-weight: bold; color: #000000; background-color: #FFEB3B; text-decoration: none; border: 4px solid #000000; text-transform: uppercase;">Reset Password</a>
                        </td>
                    </tr>
                </table>
                <p style="margin: 20px 0; font-size: 14px; color: #666666;">
                    This link will expire in <strong>15 minutes</strong>.
                </p>
                <p style="margin: 20px 0; font-size: 14px; color: #666666;">
                    If you didn't request a password reset, please ignore this email or contact support if you have concerns.
                </p>
                <hr style="border: none; border-top: 2px solid #000000; margin: 30px 0;">
                <p style="margin: 0; font-size: 12px; color: #999999;">
                    If the button doesn't work, copy and paste this link into your browser:<br>
                    <a href="%s" style="color: #000000;">%s</a>
                </p>
            </td>
        </tr>
        <tr>
            <td style="padding: 20px 30px; text-align: center; background-color: #000000; color: #ffffff; font-size: 12px;">
                &copy; 2026 AceThatPaper. All rights reserved.
            </td>
        </tr>
    </table>
</body>
</html>
`, resetLink, resetLink, resetLink)

	// Plain text alternative
	plainBody := fmt.Sprintf(`
Password Reset Request

We received a request to reset your AceThatPaper password.

Click the link below to reset your password:
%s

This link will expire in 15 minutes.

If you didn't request a password reset, please ignore this email.

© 2026 AceThatPaper. All rights reserved.
`, resetLink)

	m.SetBodyString(mail.TypeTextPlain, plainBody)
	m.AddAlternativeString(mail.TypeTextHTML, htmlBody)

	// Create the SMTP client
	// Port 587: Use STARTTLS
	// Port 465: Use implicit SSL/TLS
	var c *mail.Client
	var err error
	if s.smtpPort == 465 {
		// Port 465 uses implicit SSL
		c, err = mail.NewClient(s.smtpHost,
			mail.WithPort(s.smtpPort),
			mail.WithSMTPAuth(mail.SMTPAuthPlain),
			mail.WithUsername(s.smtpUser),
			mail.WithPassword(s.smtpPass),
			mail.WithSSL(),
			mail.WithTimeout(30*time.Second),
		)
	} else {
		// Port 587 uses STARTTLS
		c, err = mail.NewClient(s.smtpHost,
			mail.WithPort(s.smtpPort),
			mail.WithSMTPAuth(mail.SMTPAuthPlain),
			mail.WithUsername(s.smtpUser),
			mail.WithPassword(s.smtpPass),
			mail.WithTLSPortPolicy(mail.TLSOpportunistic),
			mail.WithTimeout(30*time.Second),
		)
	}
	if err != nil {
		s.logger.Println("error creating mail client:", err)
		return err
	}

	s.logger.Printf("attempting to send email via %s:%d", s.smtpHost, s.smtpPort)
	if err := c.DialAndSend(m); err != nil {
		s.logger.Printf("error sending email (host=%s, port=%d): %v", s.smtpHost, s.smtpPort, err)
		return err
	}

	s.logger.Printf("password reset email sent to %s", email[:3]+"***")
	return nil
}
