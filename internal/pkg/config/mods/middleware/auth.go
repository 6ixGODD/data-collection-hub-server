package middleware

type AuthConfig struct {
	SkippedPathPrefixes []string `mapstructure:"skipped_path_prefixes" yaml:"skipped_path_prefixes" default:"['/api/v1/auth/login', '/api/v1/auth/refresh', '/api/v1/auth/logout', '/ping']"`
}
