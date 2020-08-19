package main

import (
	"github.com/electricbubble/guia2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// driver, err := guia2.NewDriver(guia2.NewEmptyCapabilities(), "http://localhost:6790/wd/hub")
	driver, err := guia2.NewDriver(nil, "http://192.168.1.28:6790/wd/hub")
	checkErr(err)

	// fmt.Println(driver.Source())
	// return

	deviceSize, err := driver.DeviceSize()
	checkErr(err)

	var startX, startY, endX, endY int
	startX = deviceSize.Width / 2
	startY = deviceSize.Height / 2
	endX = startX
	endY = startY / 2
	err = driver.Swipe(startX, startY, endX, endY)
	checkErr(err)

	var startPoint, endPoint guia2.PointF
	startPoint = guia2.PointF{X: float64(startX), Y: float64(startY)}
	endPoint = guia2.PointF{X: startPoint.X, Y: startPoint.Y * 1.6}
	err = driver.SwipePointF(startPoint, endPoint)
	checkErr(err)

	element, err := driver.FindElement(guia2.BySelector{ResourceIdID: "tv.danmaku.bili:id/expand_search"})
	checkErr(err)

	err = element.Click()
	checkErr(err)

	bySelector := guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().Focused(true).String()}
	exists := func(d *guia2.Driver) (bool, error) {
		element, err = d.FindElement(bySelector)
		if err == nil {
			return true, nil
		}
		return false, nil
	}

	err = driver.Wait(exists)
	checkErr(err)

	err = element.SendKeys("雾山五行")
	checkErr(err)

	err = driver.PressKeyCode(guia2.KCEnter, guia2.KMEmpty)
	checkErr(err)

	bySelector = guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().TextStartsWith("番剧").String()}
	checkErr(driver.Wait(exists))
	checkErr(element.Click())

	bySelector = guia2.BySelector{UiAutomator: guia2.NewUiSelectorHelper().Text("立即观看").String()}
	checkErr(driver.Wait(exists))
	checkErr(element.Click())

	bySelector = guia2.BySelector{ResourceIdID: "tv.danmaku.bili:id/videoview_container_space"}
	checkErr(driver.Wait(exists))

	// time.Sleep(time.Second * 5)

	screenshot, err := element.Screenshot()
	checkErr(err)
	userHomeDir, _ := os.UserHomeDir()
	checkErr(ioutil.WriteFile(userHomeDir+"/Desktop/element.png", screenshot.Bytes(), 0600))

	err = driver.PressKeyCode(guia2.KCMediaPause, guia2.KMEmpty)
	checkErr(err)

	err = driver.PressBack()
	checkErr(err)
}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
