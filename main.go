package main

func main() {
	sq := square{
		sideLength: 10.0,
	}

	tr := triangle{
		height: 10.0,
		base:   10.0,
	}

	printArea(sq)
	printArea(tr)
}
