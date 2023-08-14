package comms

import (
	"fmt"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/utils"
	"time"
	"tinygo.org/x/bluetooth"
)

type Packet struct {
	bytes []byte
}

func NewPacket(flags byte, TID, SID, DID, CID, seqNum uint8, data []byte) Packet {
	bytes := []byte{0x0, flags}

	if TID != 0 {
		bytes = append(bytes, TID)
	}
	if SID != 0 {
		bytes = append(bytes, SID)
	}
	if DID != 0 {
		bytes = append(bytes, DID)
	}
	if CID != 0 {
		bytes = append(bytes, CID)
	}

	bytes = append(bytes, seqNum)
	bytes = append(bytes, data...)

	var checkSum uint8
	for _, b := range bytes {
		checkSum += b
	}
	bytes = append(bytes, ^checkSum)
	bytes = append(bytes, 0xd8)
	bytes[0] = 0x8d

	return Packet{bytes: bytes}
}

func (packet *Packet) Send(api bluetooth.DeviceCharacteristic) error {
	size, err := api.WriteWithoutResponse(packet.bytes)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Sent %02d bytes: %s\n", size, utils.ByteString(packet.bytes))
	return nil
}
