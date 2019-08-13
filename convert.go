package kata

import "fmt"

func convertToDigitalNumber(number int) int {
	digitalNumber := [8]string{
		"   ", "|  ", "  |", "| |",
		"  _ ", "|_ ", " _|", "|_|",
	}
	fmt.Println(digitalNumber[0]+"\n", digitalNumber[1]+"\n", digitalNumber[1]+"\n")
	return 1
}
