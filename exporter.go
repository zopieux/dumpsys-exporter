package main

import (
	"fmt"
	"log"
	"time"

	"github.com/electricbubble/gadb"
	"google.golang.org/protobuf/proto"

	"github.com/zopieux/dumpsys-exporter/proto/android/frameworks/base/core/proto/android/server"
	"github.com/zopieux/dumpsys-exporter/proto/android/frameworks/base/core/proto/android/service"
)

func getPower(dev gadb.Device) *server.PowerManagerServiceDumpProto {
	out, err := dev.RunShellCommandWithBytes("dumpsys power --proto")
	if err != nil {
		return nil
	}
	msg := &server.PowerManagerServiceDumpProto{}
	if err := proto.Unmarshal(out, msg); err != nil {
		return nil
	}
	return msg
}

func getProcessStats(dev gadb.Device) *service.ProcessStatsServiceDumpProto {
	out, err := dev.RunShellCommandWithBytes("dumpsys procstats --proto")
	if err != nil {
		return nil
	}
	msg := &service.ProcessStatsServiceDumpProto{}
	if err := proto.Unmarshal(out, msg); err != nil {
		return nil
	}
	return msg
}

func main() {
	adbClient, err := gadb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", adbClient)
	if err = adbClient.Connect("192.168.1.50", 5555); err != nil {
		log.Fatal(err)
	}
	c, err := adbClient.DeviceList()
	if err != nil {
		log.Fatal(err)
	}
	dev := c[0]
	model, err := dev.Model()
	fmt.Printf("%s\n", model)

	for {
		power := getPower(dev)
		fmt.Printf("%v %v\n", power.Wakefulness, power.IsDeviceIdleMode)

		stat := getProcessStats(dev)
		fmt.Printf("%+v\n", stat.ProcstatsNow.ProcessStats[0].States)

		time.Sleep(time.Second * 1)
	}
}
