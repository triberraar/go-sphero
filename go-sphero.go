package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/sphero"
)

type color struct {
	r uint8
	g uint8
	b uint8
}

const blat string = "sdf"

func main() {
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the sphero port: ")
	spheroPort, _ := reader.ReadString('\n')
	// spheroPort = "/dev/tty.Sphero-GRY-AMP-SPP"
	spheroPort = strings.TrimSpace(spheroPort)

	adaptor := sphero.NewAdaptor(spheroPort)
	driver := sphero.NewSpheroDriver(adaptor)

	wg.Add(1)
	go goSphero(driver, adaptor, wg)
	wg.Wait()
}

func goSphero(driver *sphero.SpheroDriver, adaptor *sphero.Adaptor, wg sync.WaitGroup) {
	red := color{255, 0, 0}
	green := color{0, 255, 0}
	blue := color{0, 0, 255}
	yellow := color{255, 255, 0}
	work := func() {
		defer wg.Done()
		fmt.Print("Colors used are: ")
		fmt.Print("Red")
		driver.SetRGB(red.r, red.g, red.b)
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Green")
		driver.SetRGB(green.r, green.g, green.b)
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Blue")
		driver.SetRGB(blue.r, blue.g, blue.b)
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Yellow")
		driver.SetRGB(yellow.r, yellow.g, yellow.b)
		time.Sleep(1000 * time.Millisecond)
		driver.Stop()
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	error := robot.Start()
	if error != nil {
		fmt.Print("something went horribly wrong")
		log.Fatal(error)
	}
	fmt.Print("done")
}
