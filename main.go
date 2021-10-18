package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pterm/pterm"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

const Interval = 500 * time.Millisecond

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Bli", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("nk", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	// load all the drivers:
	if _, err := host.Init(); err != nil {
		pterm.Error.Printf("host initiation failed %s\n", err)

		return
	}

	t := time.NewTicker(Interval)
	defer t.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for l := gpio.Low; ; l = !l {
		// lookup a pin by its location on the board:
		if err := gpioreg.ByName("26").Out(l); err != nil {
			pterm.Error.Printf("pinout %s\n", err)

			return
		}

		select {
		case <-t.C:
		case <-quit:
			return
		}
	}
}
