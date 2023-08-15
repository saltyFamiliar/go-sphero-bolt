package main

import (
	"time"

	"github.com/saltyFamiliar/go-sphero-bolt/pkg/sphero"
	"github.com/saltyFamiliar/go-sphero-bolt/pkg/utils"
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

	utils.Must("set speed", bolt.SetSpeed(225))
	time.Sleep(1250 * time.Millisecond)

	utils.Must("stop", bolt.SetSpeed(0))
	time.Sleep(1 * time.Second)

	utils.Must("turn around", bolt.SetDirection(270))

	utils.Must("set speed", bolt.SetSpeed(200))
	time.Sleep(7 * time.Second)

	utils.Must("stop", bolt.SetSpeed(0))
	time.Sleep(1 * time.Second)

	utils.Must("turn around", bolt.SetDirection(90))

	utils.Must("set speed", bolt.SetSpeed(200))
	time.Sleep(7 * time.Second)

	utils.Must("turn around", bolt.SetDirection(180))
	utils.Must("set speed", bolt.SetSpeed(200))
	time.Sleep(1 * time.Second)

	utils.Must("stop", bolt.SetSpeed(0))
	time.Sleep(1 * time.Second)

	utils.Must("turn back around", bolt.Rotate(0))

	// roll along square path (rhomboid in practice)
	//for i := 0; i < 4; i++ {
	//	utils.Must("set direction", bolt.SetDirection(i*90))

	//	utils.Must("start moving", bolt.SetSpeed(20))

	//	time.Sleep(2 * time.Second)

	//	utils.Must("stop moving", bolt.SetSpeed(0))
	//}

	matrix := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
	}

	bolt.SetLEDMatrix(matrix, 255, 0, 0)
	time.Sleep(3 * time.Second)

	utils.Must("turn off", bolt.PowerOff())
	time.Sleep(1 * time.Second)
}
