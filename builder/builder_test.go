package creational

import "testing"


//type ManufacturingDirector struct{}

//func (f *ManufacturingDirector) Construct() {
	//Implementation goes here
//}
//func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	//Implementation goes here
//}



func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and the were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be a 'Car' and was %s\n", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1
	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}
	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be a 'Motorbike' and was %s\n", motorbike.Structure)
	}
}
