package goodRemoteControl

import (
	"fmt"
)

type TV struct{}

func (t TV) On() {
	fmt.Println("TV를 켭니다.")
}

func (t TV) Off() {
	fmt.Println("TV를 끕니다.")
}

func (t TV) SetInputChannel(channel int) {
	fmt.Printf("채널을 %d로 변경합니다.\n", channel)
}

func (t TV) SetVolume(volume int) {
	fmt.Printf("볼륨을 %d로 변경합니다.\n", volume)
}

type AirConditioner struct{}

func (a AirConditioner) On() {
	fmt.Println("에어컨을 켭니다.")
}

func (a AirConditioner) Off() {
	fmt.Println("에어컨을 끕니다.")
}

func (a AirConditioner) SetTemperature(temperature int) {
	fmt.Printf("온도를 %d로 변경합니다.\n", temperature)
}

type Light struct{}

func (l Light) On() {
	fmt.Println("전구를 켭니다.")
}

func (l Light) Off() {
	fmt.Println("전구를 끕니다.")
}

type Window struct{}

func (w Window) Open() {
	fmt.Println("창문을 엽니다.")
}

func (w Window) Close() {
	fmt.Println("창문을 닫습니다.")
}

type Command interface {
	Execute()
}

type NoCommand struct{}

func (n NoCommand) Execute() {}

type TVOnCommand struct {
	tv TV
}

func (t *TVOnCommand) TVOnCommand(tv TV) {
	t.tv = tv
}

func (t TVOnCommand) Execute() {
	t.tv.On()
}

type TVOffCommand struct {
	tv TV
}

func (t *TVOffCommand) TVOffCommand(tv TV) {
	t.tv = tv
}

func (t TVOffCommand) Execute() {
	t.tv.Off()
}

type TVSetInputChannelCommand struct {
	tv      TV
	channel int
}

func (t *TVSetInputChannelCommand) TVSetInputChannelCommand(tv TV, channel int) {
	t.tv = tv
	t.channel = channel
}

func (t TVSetInputChannelCommand) Execute() {
	t.tv.SetInputChannel(t.channel)
}

type TVSetVolumeCommand struct {
	tv     TV
	volume int
}

func (t *TVSetVolumeCommand) TVSetVolumeCommand(tv TV, volume int) {
	t.tv = tv
	t.volume = volume
}

func (t TVSetVolumeCommand) Execute() {
	t.tv.SetVolume(t.volume)
}

type AirConditionerOnCommand struct {
	ac AirConditioner
}

func (a *AirConditionerOnCommand) AirConditionerOnCommand(ac AirConditioner) {
	a.ac = ac
}

func (a AirConditionerOnCommand) Execute() {
	a.ac.On()
}

type AirConditionerOffCommand struct {
	ac AirConditioner
}

func (a *AirConditionerOffCommand) AirConditionerOffCommand(ac AirConditioner) {
	a.ac = ac
}

func (a AirConditionerOffCommand) Execute() {
	a.ac.Off()
}

type AirConditionerSetTemperatureCommand struct {
	ac          AirConditioner
	temperature int
}

func (a *AirConditionerSetTemperatureCommand) AirConditionerSetTemperatureCommand(ac AirConditioner, temperature int) {
	a.ac = ac
	a.temperature = temperature
}

func (a AirConditionerSetTemperatureCommand) Execute() {
	a.ac.SetTemperature(a.temperature)
}

type LightOnCommand struct {
	light Light
}

func (l *LightOnCommand) LightOnCommand(light Light) {
	l.light = light
}

func (l LightOnCommand) Execute() {
	l.light.On()
}

type LightOffCommand struct {
	light Light
}

func (l *LightOffCommand) LightOffCommand(light Light) {
	l.light = light
}

func (l LightOffCommand) Execute() {
	l.light.Off()
}

type WindowOpenCommand struct {
	window Window
}

func (w *WindowOpenCommand) WindowOpenCommand(window Window) {
	w.window = window
}

func (w WindowOpenCommand) Execute() {
	w.window.Open()
}

type WindowCloseCommand struct {
	window Window
}

func (w *WindowCloseCommand) WindowCloseCommand(window Window) {
	w.window = window
}

func (w WindowCloseCommand) Execute() {
	w.window.Close()
}

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
}

func (r *RemoteControl) RemoteControl(commandCount ...int) {
	count := 5
	if len(commandCount) > 0 {
		count = commandCount[0]
	}

	r.onCommands = make([]Command, count)
	r.offCommands = make([]Command, count)

	for i := 0; i < count; i++ {
		r.onCommands[i] = NoCommand{}
		r.offCommands[i] = NoCommand{}
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r RemoteControl) OnButtonWasPushed(slot int) {
	r.onCommands[slot].Execute()
}

func (r RemoteControl) OffButtonWasPushed(slot int) {
	r.offCommands[slot].Execute()
}

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) SetCommands(commands []Command) {
	m.commands = commands
}

func (m MacroCommand) Execute() {
	fmt.Println("MacroCommand 실행")
	for _, command := range m.commands {
		(command).Execute()
	}
}

func Simulate() {
	remoteControl := RemoteControl{}
	remoteControl.RemoteControl(10)

	tv := TV{}
	tvOnCommand := TVOnCommand{tv}
	tvOffCommand := TVOffCommand{tv}
	tvSetInputChannelCommand := TVSetInputChannelCommand{tv, 3}
	tvSetVolumeCommand := TVSetVolumeCommand{tv, 10}

	airConditioner := AirConditioner{}
	airConditionerOnCommand := AirConditionerOnCommand{airConditioner}
	airConditionerOffCommand := AirConditionerOffCommand{airConditioner}
	airConditionerSetTemperatureCommand := AirConditionerSetTemperatureCommand{airConditioner, 20}

	light := Light{}
	lightOnCommand := LightOnCommand{light}
	lightOffCommand := LightOffCommand{light}

	window := Window{}
	windowOpenCommand := WindowOpenCommand{window}
	windowCloseCommand := WindowCloseCommand{window}

	remoteControl.SetCommand(0, tvOnCommand, tvOffCommand)
	remoteControl.SetCommand(1, tvSetInputChannelCommand, NoCommand{})
	remoteControl.SetCommand(2, tvSetVolumeCommand, NoCommand{})
	remoteControl.SetCommand(3, airConditionerOnCommand, airConditionerOffCommand)
	remoteControl.SetCommand(4, airConditionerSetTemperatureCommand, NoCommand{})

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
	remoteControl.OnButtonWasPushed(4)
	remoteControl.OffButtonWasPushed(4)

	remoteControl.SetCommand(5, lightOnCommand, lightOffCommand)
	remoteControl.SetCommand(6, windowOpenCommand, windowCloseCommand)

	remoteControl.OnButtonWasPushed(5)
	remoteControl.OffButtonWasPushed(5)
	remoteControl.OnButtonWasPushed(6)
	remoteControl.OffButtonWasPushed(6)

	macroCommand := MacroCommand{}
	macroCommand.SetCommands([]Command{tvOnCommand, airConditionerOnCommand, lightOnCommand, windowOpenCommand})

	// macroCommand가 포인터 리시버 사용시 인터페이스로 호출이 불가능하다.
	remoteControl.SetCommand(7, macroCommand, NoCommand{})
	remoteControl.OnButtonWasPushed(7)
}
