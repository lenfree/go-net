package main

import (
        "fmt"
        "os"
        "time"

        "github.com/fatih/color"
        "github.com/sparrc/go-ping"
        "gopkg.in/urfave/cli.v1"
        "net"
)

func IcmpPingCli() cli.Command {
        command := cli.Command{
                Name:      "ping",
                ShortName: "p",
                Usage:     "Ping Google.com which is the most reliable service in the Internet so far",
                Action:    pingGoogle,
        }
        return command
}

func pingGoogle(ctx *cli.Context) {
        green := color.New(color.FgGreen).SprintFunc()
        red := color.New(color.FgRed).SprintFunc()
        pinger, err := ping.NewPinger(googleAddr())
        if err != nil {
                fmt.Printf("%s", err.Error())
        }
        pinger.Count = 3
        pinger.Timeout = time.Second * 5
        pinger.Run()
        stats := pinger.Statistics()
        if stats.PacketsRecv == 0 && stats.PacketLoss == 100 {
                fmt.Printf("%s %s\n",
                  red("Please make sure you are properly connected to a network and re-run"),
                  green("go-net ip"),
                )
                os.Exit(0)
        } else {
                fmt.Printf("%s %s\n",
                  green("I can ping Google DNS address"),
                  red(stats.Addr),
                )
        }
        _, err = net.LookupHost(googleWWW())
        if err != nil {
                fmt.Printf("%s %s\n",
                  red("Resolver configured is not working, try Google DNS server at 8.8.8.8"),
                  red("and re-run go-net ping"),
                )
        } else {
                fmt.Printf("%s %s\n", green("I can resolve"), red(googleWWW()))
        }
}

func googleAddr() string {
        return "8.8.8.8"
}

func googleWWW() string {
        return "www.google.com"
}
