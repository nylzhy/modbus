package modbus

// RequestIns Request Instructions to make a ProtocolDataUnit (PDU) for client(TCP) or Master(RTU) with Modbus Mode
// 寄存器类型			功能			功能码		寄存器地址范围		读取/写入寄存器数量			操作数据
// 线圈				读取线圈			0x01		0x0000-0xFFFF		0-2000(0x001~0x7D0)		bit
// 离散输入			读取离散输入	 	 0x02		 0x0000-0xFFFF		 1-2000(0x001~0x7D0)	 bit
// 保持寄存器		读取保持寄存器		 0x03		 0x0000-0xFFFF		 1-125(0x01~0x7D)		 2bytes
// 输入寄存器		读取输入寄存器		 0x04		0x0000-0xFFFF		 1-125(0x01~0x7D)		2bytes
// 线圈				写单个线圈			0x05		0x0000-0xFFFF		1						0xFF00/0x0000(只有这两种形式，不存在数据长度问题)
// 保持寄存器		写单个保持寄存器	 0x06		0x0000-0xFFFF        1						2bytes
// 线圈				写多个线圈			0x0F 		0x0000-0XFFFF		1-2000(0x001~0x7D0)     Bit
// 保持寄存器		写多个保持寄存器	 0X10		0x0000-0xFFFF		 1-120(0x01-0x78)		2bytes
type RequestIns struct {
	FunctionCode uint16 `json:"funccode,omitempty"`   //ModbusFunctionCode，Ref modbus.go
	StartAddr    uint16 `json:"startaddr,omitempty"`  //Modbus Regigter Start Addr, 0-0xFFFF(65535)
	DataLength   uint16 `json:"datalength,omitempty"` //Modbus Data Length, 0-0x07D0(2000) for some FC, or 0-0x7D(125) for some FC
}

// PLCRequestIns Request Instructions to make a ProtocolDataUnit (PDU) for client(TCP) or Master(RTU) with PLC mode
// PLC Mode: 0/false is Read and 1/true is write
// PLC Modbus 数据地址、数据长度与功能码关系表
// 数据读取
// 数据起始地址, 	数据长度,	 		功能码
// 0-9999,	  		1-2000,				01			// 读取离散输出
// 10001-19999,	  	1-2000,          	02			// 读取离散输入
// 40001-49999,	  	1-125,          	03			// 读取保持寄存器
// 400001-465535,	1-125,				03			// 读取保持寄存器
// 30001-39999,		1-125,				04			// 读取输入寄存器
// 20001-29999,		1-125,				03/04		// 读取浮点寄存器
// 数据写入
// 数据起始地址, 	数据长度,	 		功能码
// 0-9999,	  		1,					05			// 写入单个离散输出
// 40001-49999,	  	1,          		06			// 写入单个保持寄存器
// 400001-465535,	1,					06			// 写入单个保持寄存器
// 0-9999,	  		2-1968,				15			// 写入多个离散输出
// 40001-49999,	  	2-123,          	16			// 写入多个保持寄存器
// 400001-465535,	2-123,				16			// 写入单个保持寄存器
// 20001-29999,		1-125,				06/16		// 写入多个浮点寄存器
// 上述功能实现依赖于厂家
type PLCRequestIns struct {
	Mode       bool   `json:"mode,omitempty"`       //0/false is Read, 1/true is Write
	DataAddr   uint16 `json:"dataaddr,omitempty"`   //Data Addr, 0-0xFFFF(65535) or 0-465535
	DataLength uint16 `json:"datalength,omitempty"` //Data Length, 0-0x07D0(2000) for some FC, or 0-0x7D(125) for some FC
}
