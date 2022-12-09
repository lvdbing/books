package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

type ConfigFile struct {
	Name string
	Path string
	Type string
}

func NewSetting(configFiles []ConfigFile) (*Setting, error) {
	vp := viper.New()
	var err error

	for i, f := range configFiles {
		vp.SetConfigName(f.Name)
		vp.AddConfigPath(f.Path)
		vp.SetConfigType(f.Type)

		if i == 0 {
			err = vp.ReadInConfig()
		} else {
			err = vp.MergeInConfig()
		}
	}
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange() // 监听配置文件变化
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			s.ReloadAllSection()
		})
	}()
}

// func NewSetting() (*Setting, error) {
// 	vp := viper.New()
// 	vp.SetConfigName("config")
// 	vp.AddConfigPath("configs/")
// 	vp.SetConfigType("yaml")
// 	err := vp.ReadInConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Setting{vp: vp}, nil
// }
