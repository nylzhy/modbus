// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

/*
Package modbus provides a client for MODBUS TCP and RTU/ASCII.
*/
package modbus

import (
	"fmt"
)

// FC Function Code Defintion 功能码定义
const (
	// Bit access

	//FuncCodeReadDiscreteInputs Read input status/DI /读取输入离散量状态 0x02
	FuncCodeReadDiscreteInputs = 2
	//FuncCodeReadCoils Read coil status/DO /读取线圈状态 0x01
	FuncCodeReadCoils = 1
	//FuncCodeWriteSingleCoil Force single coil /强制写单个线圈 0x05
	FuncCodeWriteSingleCoil = 5
	//FuncCodeWriteMultipleCoils Write Multiple Coils /预置多个线圈 0x0F
	FuncCodeWriteMultipleCoils = 15

	// 16-bit access

	//FuncCodeReadInputRegisters Read input registers /读取输入寄存器 0x04
	FuncCodeReadInputRegisters = 4
	//FuncCodeReadHoldingRegisters Read holding registers /读取保持寄存器 0x03
	FuncCodeReadHoldingRegisters = 3
	//FuncCodeWriteSingleRegister Preset single register /预置单个寄存器 0x06
	FuncCodeWriteSingleRegister = 6
	//FuncCodeWriteMultipleRegisters Write Multiple registers /写多个寄存器 0x10
	FuncCodeWriteMultipleRegisters = 16
	//FuncCodeReadWriteMultipleRegisters Read/Write Multiple registers /读写多个寄存器 0x17
	FuncCodeReadWriteMultipleRegisters = 23
	// FuncCodeMaskWriteRegister Mask Write Register /屏蔽写寄存器 0x16
	FuncCodeMaskWriteRegister = 22
	//FuncCodeReadFIFOQueue Read FIFO Queue /读取文件队列 0x18
	FuncCodeReadFIFOQueue = 24
)

/*
代码	名称	含义
01	非法功能	对于服务器（或从站）来说，询问中接收到的功能码是不可允许的操作，可能是因为功能码仅适用于新设备而被选单元中不可实现同时，还指出服务器（或从站）在错误状态中处理这种请求，例如：它是未配置的，且要求返回寄存器值。
02	非法数据地址	对于服务器（或从站）来说，询问中接收的数据地址是不可允许的地址，特别是参考号和传输长度的组合是无效的。对于带有100个寄存器的控制器来说，偏移量96和长度4的请求会成功，而偏移量96和长度5的请求将产生异常码02。
03	非法数据值	对于服务器（或从站）来说，询问中包括的值是不可允许的值。该值指示了组合请求剩余结构中的故障。例如：隐含长度是不正确的。modbus协议不知道任何特殊寄存器的任何特殊值的重要意义，寄存器中被提交存储的数据项有一个应用程序期望之外的值。
04	从站设备故障	当服务器（或从站）正在设法执行请求的操作时，产生不可重新获得的差错。
05	确认收到	与编程命令一起使用，服务器（或从站）已经接受请求，并且正在处理这个请求，但是需要长持续时间进行这些操作，返回这个响应防止在客户机（或主站）中发生超时错误，客户机（或主机）可以继续发送轮询程序完成报文来确认是否完成处理。
07	从属设备忙	与编程命令一起使用，服务器（或从站）正在处理长持续时间的程序命令，当服务器（或从站）空闲时，客户机（或主站）应该稍后重新传输报文。
08	存储奇偶性差错	与功能码20和21以及参考类型6一起使用，指示扩展文件区不能通过一致性校验。服务器（或从站）设备读取记录文件，但在存储器中发现一个奇偶校验错误。客户机（或主机）可重新发送请求，但可以在服务器（或从站）设备上要求服务。
0A	不可用网关路径	与网关一起使用，指示网关不能为处理请求分配输入端口值输出端口的内部通信路径，通常意味着网关是错误配置的或过载的。
0B	网关目标设备响应失败	与网关一起使用，指示没有从目标设备中获得响应，通常意味着设备未在网络中
*/

const (
	//ExceptionCodeIllegalFunction ILLEGAL FUNCTION 非法功能
	ExceptionCodeIllegalFunction = 1
	//ExceptionCodeIllegalDataAddress ILLEGAL DATA ADDRESS 非法数据地址
	ExceptionCodeIllegalDataAddress = 2
	//ExceptionCodeIllegalDataValue ILLEGAL DATA VALUE 非法数据值
	ExceptionCodeIllegalDataValue = 3
	//ExceptionCodeServerDeviceFailure SERVER DEVICE FAILURE 从站设备故障
	ExceptionCodeServerDeviceFailure = 4
	//ExceptionCodeAcknowledge ACKNOWLEDGE 确认
	ExceptionCodeAcknowledge = 5
	//ExceptionCodeServerDeviceBusy SERVER DEVICE BUSY 从属设备忙
	ExceptionCodeServerDeviceBusy = 6
	//ExceptionCodeMemoryParityError MEMORY PARITY ERROR  存储奇偶性差错
	ExceptionCodeMemoryParityError = 8
	//ExceptionCodeGatewayPathUnavailable GATEWAY PATH UNAVAILABLE 不可用网关路径
	ExceptionCodeGatewayPathUnavailable = 10
	//ExceptionCodeGatewayTargetDeviceFailedToRespond GATEWAY TARGET DEVICE FAILED TO RESPOND 网关目标设备响应失败
	ExceptionCodeGatewayTargetDeviceFailedToRespond = 11
)

// MError implements error interface.
type MError struct {
	FunctionCode  byte
	ExceptionCode byte
}

// Error converts known modbus exception code to error message.
func (e *MError) Error() string {
	var name string
	switch e.ExceptionCode {
	case ExceptionCodeIllegalFunction:
		name = "ILLEGAL FUNCTION/非法功能"
	case ExceptionCodeIllegalDataAddress:
		name = "ILLEGAL DATA ADDRESS/非法数据地址"
	case ExceptionCodeIllegalDataValue:
		name = "ILLEGAL DATA VALUE/非法数据值"
	case ExceptionCodeServerDeviceFailure:
		name = "SERVER DEVICE FAILURE/从站设备故障"
	case ExceptionCodeAcknowledge:
		name = "ACKNOWLEDGE/确认"
	case ExceptionCodeServerDeviceBusy:
		name = "SERVER DEVICE BUSY/从属设备忙"
	case ExceptionCodeMemoryParityError:
		name = "MEMORY PARITY ERROR/存储奇偶性差错"
	case ExceptionCodeGatewayPathUnavailable:
		name = "GATEWAY PATH UNAVAILABLE/不可用网关路径"
	case ExceptionCodeGatewayTargetDeviceFailedToRespond:
		name = "GATEWAY TARGET DEVICE FAILED TO RESPOND/网关目标设备响应失败"
	default:
		name = "UNKNOWN/未知错误"
	}
	return fmt.Sprintf("modbus: exception '%v' (%s), function '%v'", e.ExceptionCode, name, e.FunctionCode)
}

// ProtocolDataUnit (PDU) is independent of underlying communication layers.
type ProtocolDataUnit struct {
	FunctionCode byte
	Data         []byte
}

// Packager specifies the communication layer.
type Packager interface {
	Encode(pdu *ProtocolDataUnit) (adu []byte, err error)
	Decode(adu []byte) (pdu *ProtocolDataUnit, err error)
	Verify(aduRequest []byte, aduResponse []byte) (err error)
}

// Transporter specifies the transport layer.
type Transporter interface {
	Send(aduRequest []byte) (aduResponse []byte, err error)
}

//InsStru modbus instruction structure with normal style of pdu
type InsStru struct {
	FunctionCode uint16 `json:"function_code,omitempty"`
	RegAddr      uint16 `json:"reg_addr,omitempty"`
	Length       uint16 `json:"length,omitempty"`
	DataBuf      []byte `json:"data_buf,omitempty"`
}

//RInsStru modbus read instruction structure with normal style of pdu
type RInsStru struct {
	FunctionCode uint16 `json:"function_code,omitempty"`
	RegAddr      uint16 `json:"reg_addr,omitempty"`
	Length       uint16 `json:"length,omitempty"`
}

//PLCInsStru modbus instruction structure with PLC style of pdu
type PLCInsStru struct {
	RWMode  bool   `json:"rw_mode"`
	RegAddr uint   `json:"reg_addr,omitempty"`
	Length  uint16 `json:"length,omitempty"`
	DataBuf []byte `json:"data_buf,omitempty"`
}

//PLCRInsStru modbus read instruction structure with PLC style of pdu
type PLCRInsStru struct {
	RegAddr uint   `json:"reg_addr,omitempty"`
	Length  uint16 `json:"length,omitempty"`
}
