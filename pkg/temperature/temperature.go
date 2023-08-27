package temperature

import "os"

func Get(path string, first chan []byte, second chan byte) {
	data := make([]byte, 6)
	sensor, _ := os.Open(path)
	sensor.Read(data)
	first <- data[:2]
	second <- data[2]
}
