package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"Gibbon/cfg"
	"Gibbon/locales"
	"Gibbon/repl"
)

func main() {
	L, err := locales.New(locales.EnUs)
	if err != nil {
		panic(err)
	}

	host, _ := os.Hostname()

	fmt.Println(L.F("REPL_INIT", map[string]string{
		"version":  cfg.VERSION,
		"date":     time.Now().Format("January 02 2006, 15:04:05"),
		"hostname": host,
		"platform": runtime.GOOS,
	}))

	replD := repl.New(os.Stdin, os.Stdout)

	replD.Start()
}
