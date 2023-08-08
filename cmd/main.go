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
	//Must("turn on grid", bolt.LightUpGrid())
	//time.Sleep(1 * time.Second)
	Must("roll", bolt.Roll())
	time.Sleep(1 * time.Second)
	Must("roll", bolt.Roll())
	time.Sleep(1 * time.Second)
	//Must("stop rolling", bolt.StopRoll())
	//time.Sleep(1 * time.Second)
	Must("turn off", bolt.PowerOff())
	time.Sleep(1 * time.Second)
}
