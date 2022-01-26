package sample

import (
	"github.com/google/uuid"
	"math/rand"
	"pcbook-go/pb"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
func randomCPUName(brand string) string {
	if brand == "intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 7 PRO  3600U",
		"Ryzen 7 PRO 49990U",
	)
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 3080",
			"RTX 2070-SUPER",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX Vega-56",
	)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution

}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}
