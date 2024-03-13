package route

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
	DisableRegistration() bool
	ReverseProxyAuthentication() bool
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
	StorageFolder() string
}
