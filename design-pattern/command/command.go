package main

// Command 命令抽象接口
type Command interface {
	Execute() // 执行命令的方法
	Undo()    // 撤销命令的方法
}

type LightCommand struct {
	light *Light
	isOn  bool
}

func NewLightCommand(light *Light) *LightCommand {
	return &LightCommand{light: light}
}

func (l *LightCommand) Execute() {
	l.switchOn()
}

func (l *LightCommand) Undo() {
	l.switchOn()
}

func (l *LightCommand) switchOn() {
	if l.isOn {
		l.light.Off()
	} else {
		l.light.On()
	}
	l.isOn = !l.isOn
}

type StereoCommand struct {
	stereo *Stereo
}

func NewStereoCommand(stereo *Stereo) *StereoCommand {
	return &StereoCommand{stereo: stereo}
}

func (s *StereoCommand) Execute() {
	s.stereo.On()
	s.stereo.SetCD()
	s.stereo.SetVolume(11)
}

func (s *StereoCommand) Undo() {
	s.stereo.Off()
}

type FanCommand struct {
	fan       *Fan
	prevSpeed int
}

func NewFanCommand(fan *Fan) *FanCommand {
	return &FanCommand{fan: fan}
}

func (f *FanCommand) Execute() {
	f.prevSpeed = f.fan.GetSpeed()
	f.fan.NextSpeed()
}

func (f *FanCommand) Undo() {
	switch f.prevSpeed {
	case SPEED_HIGH:
		f.fan.High()
	case SPEED_MEDIUM:
		f.fan.Medium()
	case SPEED_LOW:
		f.fan.Low()
	case SPEED_OFF:
		f.fan.Off()
	}
}
