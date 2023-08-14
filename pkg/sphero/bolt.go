package sphero

import (
	"fmt"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/comms"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/flag"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/utils"
	"tinygo.org/x/bluetooth"
)

type SpheroBolt struct {
	Api         bluetooth.DeviceCharacteristic
	Connection  *bluetooth.Device
	seq         uint8
	orientation uint16
}

func NewBolt(adapter *bluetooth.Adapter, name string) (*SpheroBolt, error) {
	var boltAddr bluetooth.Address
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		if device.LocalName() == name {
			boltAddr = device.Address
			if err := adapter.StopScan(); err != nil {
				panic(nil)
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't find bolt: %s\n", err.Error())
	}

	dev, err := adapter.Connect(boltAddr, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to bolt: %s\n", err.Error())
	}

	services, err := dev.DiscoverServices(nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't discover services: %s\n", err.Error())
	}

	var apiService bluetooth.DeviceService
	for _, s := range services {
		if s.String()[:8] == "00010001" {
			apiService = s
			break
		}
	}

	characteristics, err := apiService.DiscoverCharacteristics(nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't discover characteristics: %s\n", err.Error())
	}

	var apiCh bluetooth.DeviceCharacteristic
	for _, ch := range characteristics {
		if ch.String()[:8] == "00010002" {
			apiCh = ch
			break
		}
	}

	var notifyBuf []byte
	err = apiCh.EnableNotifications(func(buf []byte) {
		notifyBuf = append(notifyBuf, buf[0])
		if buf[0] == 0xd8 {
			fmt.Println("full notification: ", utils.ByteString(notifyBuf))
			notifyBuf = []byte{}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't enable notifications: %s\n", err.Error())
	}

	return &SpheroBolt{Api: apiCh, Connection: dev}, nil
}

func (bot *SpheroBolt) NextSeq() uint8 {
	bot.seq += 1
	return bot.seq - 1
}

func (bot *SpheroBolt) PowerOn() error {
	packet := comms.NewPacket(
		flag.IsActivity|flag.RequestResponse,
		0x0,
		0x0,
		0x13,
		0x0d,
		bot.NextSeq(),
		[]byte{})
	return packet.Send(bot.Api)
}

func (bot *SpheroBolt) SetSpeed(speed uint8) error {
	highByte := uint8(bot.orientation >> 8)
	lowByte := uint8(bot.orientation & 0xFF)
	packet := comms.NewPacket(
		flag.IsActivity|flag.RequestResponse|flag.HasTargetID|flag.HasSourceID,
		0x12,
		0x01,
		0x16,
		0x07,
		bot.NextSeq(),
		[]byte{speed, highByte, lowByte, 0x00})
	return packet.Send(bot.Api)
}

func (bot *SpheroBolt) SetDirection(degrees int) error {
	degreeBytes := uint16(degrees % 360)

	highByte := uint8(degreeBytes >> 8)
	lowByte := uint8(degreeBytes & 0xFF)

	packet := comms.NewPacket(
		flag.IsActivity|flag.RequestResponse|flag.HasTargetID|flag.HasSourceID,
		0x12,
		0x01,
		0x16,
		0x07,
		bot.NextSeq(),
		[]byte{0x00, highByte, lowByte, 0x00})
	bot.orientation = degreeBytes
	return packet.Send(bot.Api)
}

func (bot *SpheroBolt) PowerOff() error {
	packet := comms.NewPacket(
		flag.IsActivity|flag.RequestResponse,
		0x0,
		0x0,
		0x13,
		0x01,
		bot.NextSeq(),
		[]byte{})
	utils.Must("send power off", packet.Send(bot.Api))
	utils.Must("disconnect", bot.Connection.Disconnect())
	return nil
}

func (bot *SpheroBolt) SetPixel(x, y, r, g, b uint8) error {
	packet := comms.NewPacket(
		0x3A,
		0x12,
		0x01,
		0x1a,
		0x2d,
		bot.NextSeq(),
		[]byte{x, y, r, g, b})
	return packet.Send(bot.Api)
}

func (bot *SpheroBolt) LightUpGrid(r, g, b uint8) error {
	packet := comms.NewPacket(
		flag.RequestResponse|flag.IsActivity|flag.HasTargetID|flag.HasSourceID,
		0x12,
		0x01,
		0x1a,
		0x2f,
		bot.NextSeq(),
		[]byte{r, g, b})
	return packet.Send(bot.Api)
}
