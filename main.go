package main

import (
	"image"
	"log"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/devices/ssd1306/image1bit"
	"periph.io/x/periph/host"
)

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
	// Draw on it.
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
			"ODROID-HC4 status",
			t.Format("Jan 2 3:04:05 PM"),
		}
		// Reset canvas
		img := image1bit.NewVerticalLSB(dev.Bounds())
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
