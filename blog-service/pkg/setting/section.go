package setting

import "time"

var sections = make(map[string]interface{})

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

type SettingEmail struct {
	Host     string
	Port     int
	Username string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return nil
		}
	}
	return nil
}
