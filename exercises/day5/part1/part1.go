package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func getLocation(path string, lower int, upper int, location *int) (err error) {
	choice := path[0]

	if len(path) == 1 {
		if choice == 'F' || choice == 'L' {
			*location = lower
		} else if choice == 'B' || choice == 'R' {
			*location = upper
		} else {
			return fmt.Errorf("Unexpected choice: %c", choice)
		}
		return
	}

	var mid float64
	mid = (float64(upper) - float64(lower)) / 2.0
	if choice == 'F' || choice == 'L' {
		a := int(math.Floor(mid))
		// fmt.Printf("%s, %d, %d\n", path[1:len(path)], lower, lower+a)
		return getLocation(path[1:len(path)], lower, lower+a, location)
	} else if choice == 'B' || choice == 'R' {
		a := int(math.Ceil(mid))
		// fmt.Printf("%s, %d, %d\n", path[1:len(path)], lower+a, upper)
		return getLocation(path[1:len(path)], lower+a, upper, location)
	} else {
		return fmt.Errorf("Unexpected choice: %c", choice)
	}
}

func main() {
	fmt.Println("Day 5 - Part 1")

	readFile, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	maxSeatID := 0

	var seatIDs []int

	for _, rowSeat := range fileTextLines {
		fmt.Println(rowSeat)
		row := rowSeat[0:7]

		var err error
		var rowLocation int
		err = getLocation(row, 0, 127, &rowLocation)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", row, rowLocation)

		seat := rowSeat[7:10]
		var seatLocation int
		err = getLocation(seat, 0, 7, &seatLocation)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", seat, seatLocation)

		seatID := rowLocation*8 + seatLocation
		fmt.Printf("SeatID: %d\n", seatID)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		fmt.Println("")

		seatIDs = append(seatIDs, seatID)
	}
	fmt.Printf("Max SeatID: %d\n", maxSeatID)
	sort.Ints(seatIDs)
	for i, seat := range seatIDs {
		if i == len(seatIDs)-1 {
			continue
		}
		if seatIDs[i+1]-seatIDs[i] > 1 {
			fmt.Printf("%d - %d\n", seat, seatIDs[i+1]-seatIDs[i])
		}
	}
	// fmt.Println(seatIDs)
}
