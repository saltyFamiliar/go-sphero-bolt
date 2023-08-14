package main

import (
	"bolt/pkg/sphero"
	. "bolt/pkg/utils"
	"time"
	"tinygo.org/x/bluetooth"
)

const boltName = "SB-9A55"

func main() {
	adapter := bluetooth.DefaultAdapter
	// Enable BLE interface.
	Must("enable BLE stack", adapter.Enable())

	bolt, err := sphero.NewBolt(adapter, boltName)
	Must("find bolt", err)

	Must("turn on", bolt.PowerOn())
	time.Sleep(1 * time.Second)

	// roll along square path (rhomboid in practice)
	for i := 0; i < 4; i++ {
		Must("set direction", bolt.SetDirection(i*90))

		Must("start moving", bolt.SetSpeed(20))
		time.Sleep(1 * time.Second)

		Must("stop moving", bolt.SetSpeed(0))
	}

	Must("turn off", bolt.PowerOff())
	time.Sleep(1 * time.Second)
}
