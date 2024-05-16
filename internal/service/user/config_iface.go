package user

type ConfigProvider interface {
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
	JWTSecret() string
}
