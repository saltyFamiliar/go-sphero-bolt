package main

import (
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/sphero"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/utils"
	"time"
	"tinygo.org/x/bluetooth"
)

const boltName = "SB-9A55"

func main() {
	adapter := bluetooth.DefaultAdapter
	// Enable BLE interface.
	utils.Must("enable BLE stack", adapter.Enable())

	bolt, err := sphero.NewBolt(adapter, boltName)
	utils.Must("find bolt", err)

	utils.Must("turn on", bolt.PowerOn())
	time.Sleep(1 * time.Second)

	// roll along square path (rhomboid in practice)
	//for i := 0; i < 4; i++ {
	//	utils.Must("set direction", bolt.SetDirection(i*90))
	//
	//	utils.Must("start moving", bolt.SetSpeed(20))
	//
	//	utils.Must("stop moving", bolt.SetSpeed(0))
	//}
	utils.Must("set pix", bolt.SetPixel(0, 0, 255, 0, 0))
	utils.Must("set pix", bolt.SetPixel(0, 1, 255, 0, 0))
	time.Sleep(3 * time.Second)

	utils.Must("turn off", bolt.PowerOff())
	time.Sleep(1 * time.Second)
}
