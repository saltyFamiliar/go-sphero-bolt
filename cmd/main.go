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

	bolt.Move(0, 200, 3000)
	time.Sleep(1 * time.Second)
	bolt.Move(180, 200, 3000)
	time.Sleep(1 * time.Second)

	utils.Must("turn back around", bolt.Rotate(0))

	matrix := [][]byte{
		{0, 1, 1, 0, 0, 1, 1, 0},
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
