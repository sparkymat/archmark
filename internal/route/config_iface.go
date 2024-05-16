package route

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
	DisableRegistration() bool
	ReverseProxyAuthentication() bool
	ProxyAuthEmailHeader() string
	ProxyAuthNameHeader() string
	StorageFolder() string
}
