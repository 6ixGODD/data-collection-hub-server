package mods

type TasksConfig struct {
	SyncLogsSpec  string `mapstructure:"sync_logs_spec" yaml:"sync_logs_spec" default:"0 0 0 * * *"`
	UpdateKeySpec string `mapstructure:"update_key_spec" yaml:"update_key_spec" default:"0 0 0 * * *"`
}
