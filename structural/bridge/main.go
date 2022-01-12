package structural

import (
	"errors"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}
type PrinterAPI1 struct{}

func (d *PrinterAPI1) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

type PrinterAPI2 struct {
	Writer io.Writer
}

func (d *PrinterAPI2) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}
type PrinterAbstraction interface {
	Print() error
}
type NormalPrinter struct{
	Msg string
	Printer PrinterAPI
}
func (c *NormalPrinter) Print() error{
	return errors.New("Not implemented yet")
}