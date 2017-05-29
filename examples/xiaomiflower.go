// +build

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

var done = make(chan struct{})

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("Scanning...")
		d.Scan([]gatt.UUID{}, true)
		return
	default:
		fmt.Printf("Stop scanning... %v \n",s)
		d.StopScanning()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalf("usage: %s [options] peripheral-id\n", os.Args[0])
	}

	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	//d.Option(gatt.LnxSetAdvertisingEnable(true)) // Can only be used with Option.

	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

  d.Init(onStateChanged)
	//d.Scan(ss []UUID, true)

	/*
	// Register handlers.
	d.Handle(
		//gatt.PeripheralDiscovered(onPeriphDiscovered),
		//gatt.PeripheralConnected(onPeriphConnected),
		//gatt.PeripheralDisconnected(onPeriphDisconnected),
	)

	d.Init(onStateChanged)
	for {
		<-time.After(10 * time.Second)
		d.Init(onStateChanged)
	}
	*/
	<-done
	fmt.Println("Done")
}
