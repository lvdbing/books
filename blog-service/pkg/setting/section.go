package setting

import "time"

type SettingServer struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type SettingApp struct {
	DefaultPagesize    int
	MaxPagesize        int
	LogSavePath        string
	LogFileName        string
	LogFileExt         string
	UploadPath         string
	UploadUrl          string
	UploadImgMaxSize   int
	UploadImgAllowExts []string
}

type SettingDatabase struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type SettingJWT struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
