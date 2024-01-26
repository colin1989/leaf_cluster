package log

import (
	"fmt"
	"github.com/name5566/leaf/log/constant"
	"os"
	"path/filepath"
	"strings"
)

const (
	pkgVersion = "0.2.0"
)

var (
	startTime string
	goVersion string
)

var (
	moduleName      string //属于哪个模块
	serviceName     string
	appName         string
	hostName        string
	buildAppVersion string
	buildUser       string
	buildHost       string
	buildTime       string
)

//
//func init() {
//	if appName == "" {
//		appName = os.Getenv(constant.EnvAppName)
//		if appName == "" {
//			appName = filepath.Base(os.Args[0])
//		}
//	}
//
//	name, err := os.Hostname()
//	if err != nil {
//		name = "unknown"
//	}
//	hostName = name
//	startTime = time.Now().Format("2006-01-02 15:04:05")
//	SetBuildTime(buildTime)
//	goVersion = runtime.Version()
//	InitEnv()
//}

func ServiceName() string {
	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
	}

	return serviceName
}

func SetServiceName(s string) {
	serviceName = s
}

func ModuleName() string {
	return moduleName
}

func SetModuleName(name string) {
	moduleName = name
}

// AppID get appID
func AppID() string {
	return appName
}

// SetAppID set appID
func SetAppID(s string) {
	appName = s
}

// AppVersion get buildAppVersion
func AppVersion() string {
	return buildAppVersion
}

func PkgVersion() string {
	return pkgVersion
}

// BuildTime get buildTime
func BuildTime() string {
	return buildTime
}

// BuildUser get buildUser
func BuildUser() string {
	return buildUser
}

// BuildHost get buildHost
func BuildHost() string {
	return buildHost
}

// SetBuildTime set buildTime
func SetBuildTime(param string) {
	buildTime = strings.Replace(param, "--", " ", 1)
}

// HostName get host name
func HostName() string {
	return hostName
}

// StartTime get start time
func StartTime() string {
	return startTime
}

// GoVersion get go version
func GoVersion() string {
	return goVersion
}

func RegistryPrefix() string {
	if appMode == "" {
		return constant.RegistryPathPrefix
	}

	return fmt.Sprintf("%s/%s", constant.RegistryPathPrefix, appMode)
}
