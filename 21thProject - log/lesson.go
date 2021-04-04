// loggingレッスン

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// logをファイルとして出力してみる
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// Fatalを使用すると、Fatalのログを出力し、そこでプログラム処理が終了する。
	log.Fatalln("error!!")

	// 上のFatal処理でプログラムが終了しているので、このPrintlnは実行されない。
	fmt.Println("ok!")
}
