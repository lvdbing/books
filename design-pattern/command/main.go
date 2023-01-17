package main

import "fmt"

func main() {
	livingRoomLight := NewLight("Living Room Light")
	kitchenLight := NewLight("Kitchen Light")
	stereo := NewStereo()
	fan := NewFan("Living Room Ceiling Fan")

	livingRoomLightCmd := NewLightCommand(livingRoomLight)
	kitchenLightCmd := NewLightCommand(kitchenLight)
	stereoCmd := NewStereoCommand(stereo)
	fanCmd := NewFanCommand(fan)

	remote := NewRemoteControl(7)

	remote.SetCommand(0, livingRoomLightCmd)
	remote.SetCommand(1, kitchenLightCmd)
	remote.SetCommand(2, stereoCmd)
	remote.SetCommand(3, fanCmd)

	fmt.Println(remote)

	remote.ButtonWasPushed(0)
	remote.ButtonWasPushed(0)
	remote.ButtonWasPushed(1)
	remote.ButtonWasPushed(1)
	remote.ButtonWasPushed(1)
	remote.ButtonWasPushed(2)
	remote.ButtonWasPushed(2)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)
	remote.ButtonWasPushed(3)

	fmt.Println()

	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
	remote.UndoButtonWasPushed()
}
