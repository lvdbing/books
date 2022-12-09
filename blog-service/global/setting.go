package global

import (
	"github.com/lvdbing/books/blog-service/pkg/logger"
	"github.com/lvdbing/books/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.SettingServer
	AppSetting      *setting.SettingApp
	DatabaseSetting *setting.SettingDatabase
	JWTSetting      *setting.SettingJWT
	EmailSetting    *setting.SettingEmail

	Logger *logger.Logger
)
