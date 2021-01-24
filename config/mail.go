package config

var MAIL = struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
}{
	HOST:     getEnv("MAIL_HOST", "smtp.gmail.com"),
	PORT:     getEnv("MAIL_PORT", "587"),
	USERNAME: getEnv("MAIL_USERNAME", ""),
	PASSWORD: getEnv("MAIL_PASSWORD", ""),
}
