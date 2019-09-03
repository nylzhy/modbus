go cmd modbus
=========

The project based on [goburrow/modbus](https://github.com/goburrow/modbus) add some new functions
for Client for some situation which need more directly construct the modbus instructions. 

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

PLC Mode:
*   ExecPLC(PLCInsStru)


Supported formats
-----------------
*   TCP
*   Serial (RTU, ASCII)

Usage
-----

// 新建一个NewClient，初始各种参数，新建InsStru和PLCInsStru结构体，实现Exec函数

```go
type InsStru struct{
	FunctionCode uint16 //Modbus function code
	RegAddr uint16	//register start addr
	Length uint16 //read/write length
	DataBuf []byte //write data buffer
}

type PLCInsStru struct {
	RWMode  bool //false means read, and true means write
	RegAddr uint //register start addr
	Length  uint16 //read/write length
	DataBuf []byte //write data buffer
}

```



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
