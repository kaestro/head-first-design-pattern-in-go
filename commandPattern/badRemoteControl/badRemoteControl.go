package badRemoteControl

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

type RemoteControl struct {
	tv             TV
	airConditioner AirConditioner
	light          Light
	window         Window
}

func (r RemoteControl) ControlTV(command string) {
	switch command {
	case "on":
		r.tv.On()
	case "off":
		r.tv.Off()
	}
}

func (r RemoteControl) ControlAirConditioner(command string) {
	switch command {
	case "on":
		r.airConditioner.On()
	case "off":
		r.airConditioner.Off()
	}
}

func (r RemoteControl) ControlLight(command string) {
	switch command {
	case "on":
		r.light.On()
	case "off":
		r.light.Off()
	}
}

func (r RemoteControl) ControlWindow(command string) {
	switch command {
	case "open":
		r.window.Open()
	case "close":
		r.window.Close()
	}
}

func Simulate() {
	remoteControl := RemoteControl{
		tv:             TV{},
		airConditioner: AirConditioner{},
		light:          Light{},
		window:         Window{},
	}

	remoteControl.ControlTV("on")
	remoteControl.ControlTV("off")

	remoteControl.ControlAirConditioner("on")
	remoteControl.ControlAirConditioner("off")

	remoteControl.ControlLight("on")
	remoteControl.ControlLight("off")

	remoteControl.ControlWindow("open")
	remoteControl.ControlWindow("close")
}
