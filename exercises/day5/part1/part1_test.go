package main

import "testing"

func TestGetLocation(t *testing.T) {
	var err error
	var rowLocation int

	err = getLocation("FBFBBFF", 0, 127, &rowLocation)
	if err != nil {
		t.Errorf("Got error")
	}
	if rowLocation != 44 {
		t.Errorf("Expected rowLocation 44, got %d", rowLocation)
	}

	var seatLocation int
	err = getLocation("RLR", 0, 7, &seatLocation)
	if err != nil {
		t.Errorf("Got error")
	}
	if seatLocation != 5 {
		t.Errorf("Expected seatLocation 5, got %d", seatLocation)
	}

	if (rowLocation*8 + seatLocation) != 357 {
		t.Errorf("Failed final check. Expected 357, got %d", rowLocation*8+seatLocation)
	}
}
