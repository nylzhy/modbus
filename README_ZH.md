go cmd modbus
=========

[English](https://github.com/nylzhy/modbus/README.md) [中文](https://github.com/nylzhy/modbus/README_ZH.md)

这个项目是基于 [goburrow/modbus](https://github.com/goburrow/modbus) 而建立的，并且为更加直接地利用modbus指令添加新的特性。当然，你仍然可以利用原来地方法或风格去使用该包，详情参见[godoc](https://godoc.org/github.com/goburrow/modbus).

支持的功能
-------------------
位访问:
*   读取离散输入状态
*   读取线圈状态
*   写入单个线圈
*   写入多个线圈

字节访问:
*   读取输入寄存器
*   读取保持寄存器
*   写单个保持寄存器
*   写多个保持寄存器
*   读写多个寄存器
*   Mask写寄存器
*   读取队列

混合访问模式:
*   Exec(InsStru)

```go
type InsStru struct{
	FunctionCode uint16 //Modbus 功能码
	RegAddr uint16	//寄存器起始地址
	Length uint16 //数据读写长度
	DataBuf []byte //读写缓存数据
}
```

带有PLC风格的混合访问模式:
*   ExecPLC(PLCInsStru)

```go
type PLCInsStru struct {
	RWMode  bool //false为读数据，true为写数据
	RegAddr uint //PLC风格的寄存器起始地址
	Length  uint16 //数据读写长度
	DataBuf []byte //读写缓存数据
}
```

支持的协议
-----------------
*   TCP
*   Serial (RTU, ASCII)

用法
-----

// 新建一个NewClient，初始各种参数，新建InsStru和PLCInsStru结构体，实现Exec函数


```go
// Modbus TCP
handler := modbus.NewTCPClientHandler("localhost:502")
handler.Timeout = 10 * time.Second
handler.SlaveId = 0xFF
handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
// Connect manually so that multiple requests are handled in one connection session
err := handler.Connect()
defer handler.Close()

client := modbus.NewClient(handler)
results, err := client.ReadDiscreteInputs(15, 2)
results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
```

```go
// Modbus RTU/ASCII
handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
handler.BaudRate = 115200
handler.DataBits = 8
handler.Parity = "N"
handler.StopBits = 1
handler.SlaveId = 1
handler.Timeout = 5 * time.Second

err := handler.Connect()
defer handler.Close()

client := modbus.NewClient(handler)
results, err := client.ReadDiscreteInputs(15, 2)
```

参考
----------
-   [Modbus 协议规范及实现指导](http://www.modbus.org/specs.php)
-   [Siemens 1200/1500 Modbus RTU](https://support.industry.siemens.com/cs/document/109477716/s7-1500-modbus-rtu%E4%BD%BF%E7%94%A8%E5%BF%AB%E9%80%9F%E5%85%A5%E9%97%A8(%E6%9B%B4%E6%96%B0%E7%89%88)?dti=0&lc=zh-CN)
