package user

type ConfigProvider interface {
	ProxyAuthEmailHeader() string
	ProxyAuthNameHeader() string
	JWTSecret() string
}
