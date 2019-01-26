package main


import (
	"fmt"
)

const CLEAR_MSB = 0x7F

type Device struct{
	Id byte
	Channels int8
}

type CommandPacket struct{
	Command byte
	Data []byte
}

func SetTarget(channel, position int16)(ret *CommandPacket){
	ret = &CommandPacket{Command:0x84}
	ret.Data = append(ret.Data, byte(channel), byte(position & 0xFF), byte(position >> 7))
	return
} 

func SetSpeed(channel, speed int16)(ret *CommandPacket){
	ret = &CommandPacket{Command:0x87}
	ret.Data = append(ret.Data, byte(channel), byte(speed & 0xFF), byte(speed >> 7))
	return
}

func SetAccelleration(channel, accel int16)(ret *CommandPacket){
	ret = &CommandPacket{Command:0x89}
	ret.Data = append(ret.Data, byte(channel), byte(accel & 0xFF), byte(accel >> 7))
	return
}

func GetPosition(channel int8)(ret *CommandPacket){
	ret = &CommandPacket{Command:0x90}
	ret.Data = append(ret.Data, byte(channel))
	return
}
	

func (d *Device) DoCommand(cmd *CommandPacket) (data []byte){
	data = append(data, 0xAA, d.Id, cmd.Command & CLEAR_MSB)
	data = append( data, cmd.Data...)
	return //naked return
}


	


func main() {
	var dev1 = Device{1, int8(6)}
	var dev2 = Device{12, int8(6)}
	
	var accels = []int16{0,200,300,1400,500,600}
	
	for i := int8(0); i < dev1.Channels; i++{
		fmt.Printf("Set Accelleration on device %v: %v\n", dev1.Id,dev1.DoCommand(SetAccelleration(int16(i), accels[i])))
	}
	fmt.Println(dev2.DoCommand(&CommandPacket{Command:0x84, Data:[]byte{4,5,6}}))
	//x := SetTarget(1,6000)
	fmt.Printf("SetSpeed: %v\n",dev2.DoCommand(SetSpeed(1,5000)))
	fmt.Println(dev2.DoCommand(SetTarget(1,6000)))
	
	fmt.Printf("Get Position of channel %v of device: %v: %v\n", 4, dev2.Id,dev2.DoCommand(GetPosition(4)))
	
	
}
