package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sparrc/go-ping"
)

var (
	count   = flag.Uint("c", 100, "Upper limit of parallel processing")
	debug   = flag.Bool("debug", false, "Print timeout ipadresses")
	timeout = flag.Uint("t", 100, "Set of ping timeout [ms]")
)

func main() {
	flag.Parse()

	num := 255
	ipList := []string{}
	missList := []string{}

	ipHeader := "169.254"

	ch := make(chan int, *count)
	wg := sync.WaitGroup{}

	for i := (0); i < num; i++ {
		for j := (1); j < num; j++ {
			ch <- 1
			wg.Add(1)
			go func(i int, j int) {
				ip := ipHeader + "." + strconv.Itoa(i) + "." + strconv.Itoa(j)
				// print(ip)
				if sendPing(ip) == true {
					ipList = append(ipList, ip)
				} else {
					if *debug == true {
						missList = append(missList, ip)
					}
				}
				<-ch
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()

	sort.Strings(ipList)
	fmt.Println("Success\n", strings.Join(ipList, "\n"))

	if *debug == true {
		sort.Strings(missList)
		fmt.Println("Miss\n", strings.Join(missList, "\n"))
	}
}

func sendPing(ipaddress string) (result bool) {
	pinger, err := ping.NewPinger(ipaddress)
	if err != nil {
		panic(err)
	}

	pinger.Count = int(*count)
	pinger.Timeout = time.Duration(*timeout) * time.Millisecond
	pinger.Run()

	if pinger.PacketsRecv == 0 {
		return false
	} else if pinger.PacketsRecv > 0 {
		return true
	} else {
		panic("err")
	}

}
