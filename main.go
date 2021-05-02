package main

import (
	"fmt"
	"image"
	"log"
	"net"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/devices/ssd1306/image1bit"
	"periph.io/x/periph/host"
)

// based on https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func formatSize(size uint64, unit uint64) string {
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, suffix := unit, 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		suffix++
	}
	value := float64(size) / float64(div)
	var fmtStr string
	if value >= 100 {
		fmtStr = "%.f%c"
	} else {
		fmtStr = "%.1f%c"
	}

	return fmt.Sprintf(fmtStr,
		value, "kMGTP"[suffix])
}

// TODO: error handling
func getMemString() string {
	v, _ := mem.VirtualMemory()

	return fmt.Sprintf("MEM: %v/%v", formatSize(v.Used, 1024), formatSize(v.Total, 1024))
}

func getCPUString() string {
	v, _ := cpu.Percent(0, false)
	l, _ := load.Avg()
	// Unfortunately, the screen just isn't wide enough to include Load15
	return fmt.Sprintf("CPU: %.f%% (%.1f %.1f)", v[0], l.Load1, l.Load5)
}

// getHDDString returns data about the biggest mounted partition.
func getHDDString() string {
	partitions, _ := disk.Partitions(false)
	biggestDiskSize := uint64(0)
	biggestDiskUsed := uint64(0)
	biggestDiskName := ""
	for _, partition := range partitions {
		d, _ := disk.Usage(partition.Mountpoint)
		if d.Total > biggestDiskSize {
			biggestDiskName = partition.Mountpoint
			biggestDiskUsed = d.Used
			biggestDiskSize = d.Total
		}
	}
	return fmt.Sprintf("%v: %v/%v", biggestDiskName, formatSize(biggestDiskUsed, 1000), formatSize(biggestDiskSize, 1000))
}

func getIPAddrString() string {
	// https://stackoverflow.com/a/37382208/3814663
	// Note that since this is UDP, no connection is actually established.
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "IP: Network down"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return fmt.Sprintf("IP: %v", localAddr.IP)
}

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	dev, err := ssd1306.NewI2C(b, &ssd1306.Opts{
		W:             128,
		H:             64,
		Rotated:       true,
		Sequential:    false,
		SwapTopBottom: false,
	})
	if err != nil {
		log.Fatalf("failed to initialize ssd1306: %v", err)
	}

	f := basicfont.Face7x13
	drawer := font.Drawer{
		Src:  &image.Uniform{image1bit.On},
		Face: f,
		Dot:  fixed.P(0, f.Height),
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		t := <-ticker.C

		lines := []string{
			t.Format("Jan 2 3:04:05 PM"),
			getIPAddrString(),
			getCPUString(),
			getMemString(),
			getHDDString(),
		}

		img := image1bit.NewVerticalLSB(dev.Bounds()) // reset canvas per frame
		drawer.Dst = img
		for i, s := range lines {
			drawer.Dot = fixed.P(0, (f.Height-1)*(i+1))
			drawer.DrawString(s)
		}

		if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
			log.Fatal(err)
		}
	}

}
