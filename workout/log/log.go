package log

import "log"

func Print(args ...any) {
	log.Print(args...)
}

func Printf(format string, args ...any) {
	log.Printf(format, args...)
}

func Println(args ...any) {
	log.Println(args...)
}

func Fatal(args ...any) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}

func Fatalln(args ...any) {
	log.Fatalln(args...)
}

func Panic(args ...any) {
	log.Panic(args...)
}

func Panicf(format string, args ...any) {
	log.Panicf(format, args...)
}

func Panicln(args ...any) {
	log.Panicln(args...)
}
