package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter
}

type fileWriter struct {
	filePath string
}

func (f *fileWriter) Write(s string) error {
	fmt.Println("file writer")
	return nil
}

type databaseWriter struct {
	host string
	port int
}

func (d *databaseWriter) Write(s string) error {
	fmt.Println("database writer")
	return nil
}

func main() {
	wc := &writerCloser{&databaseWriter{host: "", port: 1}}
	wc.Write("sss")

	wc1 := &writerCloser{MyWriter: &fileWriter{
		filePath: "",
	}}
	wc1.Write("ddd")
}
