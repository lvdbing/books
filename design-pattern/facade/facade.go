package main

// ******* 外观模式 *******
// 外观将一群对象“包装”起来以简化其接口，将客户从一个复杂的子系统中解耦。
// 实现一个外观，需要将子系统组合进外观中，然后将工作委托给子系统执行。

type HomeTheaterFacade struct {
	// 组合用到的子系统组件
	popper Popper
	light  Light
	dvd    DVD
}

func NewHomeTheaterFacade(popper Popper, light Light, dvd DVD) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		popper: popper,
		light:  light,
		dvd:    dvd,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	// 打开爆米花机，开始爆米花。
	h.popper.On()
	h.popper.Pop()

	// 将灯光调暗。
	h.light.Dim(10)

	// 打开DVD播放器，开始播放电影。
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	// 关闭爆米花机。
	h.popper.Off()

	// 打开灯。
	h.light.On()

	// 停止播放电影并弹出，然后关闭DVD播放器。
	h.dvd.Stop()
	h.dvd.Eject()
	h.dvd.Off()
}
