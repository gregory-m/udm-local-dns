package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gregory-m/udm-se-dns/api"
	"github.com/gregory-m/udm-se-dns/config"
	"github.com/gregory-m/udm-se-dns/dnsmasq"
)

const pidFileName = "/run/dnsmasq.pid"
const confFileName = "/run/dnsmasq.conf.d/my-dns.conf"
const sleepTime = 30

var client *http.Client
var configFile = flag.String("config", "unifi-dns.toml", "config file to use")

func main() {
	flag.Parse()
	conf, err := config.ReadConfig(*configFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	client, err := api.New(conf.Username, conf.Password, conf.URL, conf.Site)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	err = client.Login()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	for {

		networks, err := client.GetNetworks()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		clients, err := client.GetClients()

		records, err := dnsmasq.Gen(clients, networks, true)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		wantS := dnsmasq.Marshal(records)

		f, err := os.OpenFile(confFileName, os.O_RDWR|os.O_CREATE, 644)
		cur, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		curS := string(cur)
		if wantS != curS {

			fmt.Println("Updating")

			err = f.Truncate(0)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			f.Seek(0, 0)

			f.Write([]byte(wantS))
			f.Close()

			pidRaw, err := ioutil.ReadFile(pidFileName)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			pid, err := strconv.Atoi(strings.Trim(string(pidRaw), "\n"))
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			err = syscall.Kill(pid, syscall.SIGTERM)
			if err != nil {
				fmt.Printf("Can't kill dnsmasq process with pid %d: %s\n", pid, err)
				os.Exit(1)
			}
			time.Sleep(sleepTime * time.Second)
			continue
		}
		f.Close()

		time.Sleep(sleepTime * time.Second)
	}

}
