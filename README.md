go cmd modbus
=========

[English](https://github.com/nylzhy/modbus/README.md) [中文](https://github.com/nylzhy/modbus/README_ZH.md)

The project base on [goburrow/modbus](https://github.com/goburrow/modbus), and add some new functions
for some situation which need more directly construct the modbus instructions. 

Of course, you still can use the old methods/style as description in [godoc](https://godoc.org/github.com/goburrow/modbus). 


Supported functions
-------------------
Bit access:
*   Read Discrete Inputs
*   Read Coils
*   Write Single Coil
*   Write Multiple Coils

16-bit access:
*   Read Input Registers
*   Read Holding Registers
*   Write Single Register
*   Write Multiple Registers
*   Read/Write Multiple Registers
*   Mask Write Register
*   Read FIFO Queue

Mix Mode:
*   Exec(InsStru)


```go
type InsStru struct{
	FunctionCode uint16 //Modbus function code
	RegAddr uint16	//register starting addr
	Length uint16 //read/write length
	DataBuf []byte //write data buffer
}
```

PLC Mode:
*   ExecPLC(PLCInsStru)
```go
type PLCInsStru struct {
	RWMode  bool //false means read, and true means write
	RegAddr uint //register starting addr
	Length  uint16 //read/write length
	DataBuf []byte //write data buffer
}
```


Supported formats
-----------------
*   TCP
*   Serial (RTU, ASCII)

Usage
-----







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

References
----------
-   [Modbus Specifications and Implementation Guides](http://www.modbus.org/specs.php)
