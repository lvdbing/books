package main

import (
	"fmt"
	"reflect"
)

// RemoteControl Invoker对象（遥控器）
type RemoteControl struct {
	Commands     []Command // Invoker对象持有Command的引用
	UndoCommands []Command
}

func NewRemoteControl(len int) *RemoteControl {
	r := &RemoteControl{}
	r.Commands = make([]Command, len)

	return r
}

func (r *RemoteControl) SetCommand(slot int, command Command) {
	r.Commands[slot] = command
}

func (r *RemoteControl) ButtonWasPushed(slot int) {
	r.Commands[slot].Execute() // 调用Command的执行方法Execute
	r.UndoCommands = append(r.UndoCommands, r.Commands[slot])
}

func (r *RemoteControl) UndoButtonWasPushed() {
	if len(r.UndoCommands) == 0 {
		return
	}
	r.UndoCommands[len(r.UndoCommands)-1].Undo()
	r.UndoCommands = r.UndoCommands[:len(r.UndoCommands)-1]
}

func (r *RemoteControl) String() string {
	s := "\n------- Remote Control -------\n"
	for i := 0; i < len(r.Commands); i++ {
		s += fmt.Sprintf("[slot %d] %s\n", i,
			reflect.TypeOf(r.Commands[i]).String())
	}
	return s
}
