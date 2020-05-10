// Create a data structure to pass to a template which
// contains information about California hotels including Name, Address, City, Zip, Region
// region can be: Southern, Central, Northern
// can hold an unlimited number of hotels
package main

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type hotels []hotel

var caliHotels = hotels{
	hotel{
		"Motel 6",
		"1234 W. Hollywood Blvd",
		"Los Angeles",
		95213,
		"Southern",
	},
	hotel{
		"Days Inn",
		"9274 Columbine St",
		"Encino",
		95232,
		"Southern",
	},
}
