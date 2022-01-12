package structural

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterAPI1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
	api2 := PrinterAPI2{}
	err = api2.PrintMessage("")
}
func TestPrintAPI2(t *testing.T) {
	api2 := PrinterAPI2{}
	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterAPI2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}
	testWriter := TestWriter{}
	api2 = PrinterAPI2{
		Writer: &testWriter,
	}
	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not writer correctlyy on the io.writer \n Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
func TestNormalPrinter_Print(t *testing.T){
	expectedMessage := "Hello io.Writer"
	normal := NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterAPI1{},
	}
	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterAPI2{
			Writer: &testWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testWriter.Msg != expectedMessage{
		t.Errorf("The expected message on the io.writer doesn't match actual.\n Actual: %S\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}