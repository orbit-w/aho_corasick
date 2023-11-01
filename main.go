package main

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

/*
   @Author: orbit-w
   @File: main
   @2023 10月 周二 12:25
*/

func main() {
	PProf()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ac, _ := aho_corasick.LoadDict("./data/en/dict.txt")
	fmt.Println(ac.Cap())
	ac.Replace("outlearningsdwdsdoutgnawedsdwdsdad", '*')

	<-ch
}

func PProf() {
	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:9000")
		if err != nil {
			panic("pprof start failed")
		}

		err = http.Serve(listener, nil)
		if err != nil {
			panic("pprof start failed")
		}
	}()
}
