package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// O_RDWR: 読み書き
	// O_CREATE: 作成
	// O_APPEND: 追加
	// Permission: 0666(rw-rw-rw-) 読み書きのみ許可
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// MultiWriter: ログの書き込み先の設定
	// Stdout: 標準出力
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// SetFlags: ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    // ログの出力先
	log.SetOutput(multiLogFile)
}
