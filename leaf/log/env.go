package log

import (
	"crypto/md5"
	"fmt"
	"github.com/name5566/leaf/log/constant"
	"os"
)

var (
	appLogDir   string
	appMode     string
	appRegion   string
	appZone     string
	appHost     string
	appInstance string
	appConfAddr string
	appConfExt  string
)

func InitEnv() {
	appLogDir = os.Getenv(constant.EnvAppLogDir)
	appMode = os.Getenv(constant.EnvAppMode)
	appRegion = os.Getenv(constant.EnvAppRegion)
	appZone = os.Getenv(constant.EnvAppZone)
	appHost = os.Getenv(constant.EnvAppHost)
	appInstance = os.Getenv(constant.EnvAppInstance)
	if appInstance == "" {
		appInstance = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", HostName(), AppID()))))
	}
	appConfAddr = os.Getenv(constant.EnvAppConfAddr)
	appConfExt = os.Getenv(constant.EnvAppConfFormat)
}

func AppConfAddr() string {
	return appConfAddr
}

func SetAppConfAddr(addr string) {
	appConfAddr = addr
}

func AppConfExt() string {
	return appConfExt
}

func SetAppConfFormat(ext string) {
	appConfExt = ext
}

func AppLogDir() string {
	return appLogDir
}

func SetAppLogDir(logDir string) {
	appLogDir = logDir
}

func AppMode() string {
	return appMode
}

func SetAppMode(mode string) {
	appMode = mode
}

func AppRegion() string {
	return appRegion
}

func SetAppRegion(region string) {
	appRegion = region
}

func AppZone() string {
	return appZone
}

func SetAppZone(zone string) {
	appZone = zone
}

func AppHost() string {
	return appHost
}

func SetAppHost(host string) {
	appHost = host
}

func AppInstance() string {
	return appInstance
}
