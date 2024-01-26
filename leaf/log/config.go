package log

import (
	"strings"
)

// Config for redis, contains RedisStubConfig and RedisClusterConfig
type Config struct {
	Debug bool // 本地调试模式
	Level string

	// writer (syslog 地址)
	Net  string
	Addr string //remote addr of syslog
	Tag  string //tag of syslog

	// 日志初始化字段
	configKey   string
	serviceName string
	moduleName  string
}

func (c *Config) WithModuleName(moduleName string) *Config {
	c.moduleName = moduleName

	return c
}

func (c *Config) WithServiceName(serviceName string) *Config {
	c.serviceName = serviceName

	return c
}

func (c *Config) GetTag() string {
	if c.moduleName == "" || c.serviceName == "" {
		return c.Tag
	}

	return FirstUpper(c.moduleName) + FirstUpper(c.serviceName)
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ToLower(s)

	return strings.ToUpper(s[:1]) + s[1:]
}

func (c *Config) BuildOption() []Option {
	//
	opts := make([]Option, 0)
	// Level
	lvl, err := GetLevel(c.Level)
	if err != nil {
		lvl = InfoLevel
	}
	opts = append(opts, WithLevel(lvl))
	// CallerSkipCount 默认 3
	opts = append(opts, WithCallerSkipCount(3))

	//
	fields := make(map[string]interface{}, 0)

	fields["serviceName"] = c.serviceName
	fields["moduleName"] = c.moduleName
	opts = append(opts, WithFields(fields))

	//
	//if config.Plugins == "zap" {
	zap := ZapCnf{
		Addr:  c.Addr,
		Net:   "tcp",
		Tag:   c.GetTag(),
		Debug: c.Debug,
	}
	opts = append(opts, WithSetOption(ZapCnfKey{}, zap))
	//
	//if c.configKey != "" {
	//	global.AutoChangeLevel(c.configKey + ".level")
	//}

	return opts
}

func (c *Config) Build() {
	// todo
	lvl, err := GetLevel(c.Level)
	if err != nil {
		lvl = DebugLevel
	}
	opts := make([]Option, 0)
	opts = append(opts, WithCallerSkipCount(3), WithLevel(lvl))
	DefaultLogger = NewLogger(opts...)
	//if c.configKey != "" {
	//	global.AutoChangeLevel(c.configKey + ".level")
	//}

}
