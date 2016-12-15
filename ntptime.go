package main

import (
        "flag"
        "fmt"
        "os"
        "time"

        "github.com/bt51/ntpclient"
)

func main() {
        ip := flag.String("ip", "", "ntp service ip address")
        port := flag.Int("port", 123, "ntp service port number")
        version := flag.Bool("v", false, "Show version.")

        flag.Parse()

        //Display version info.
        if *version {
                fmt.Println("version=1.0, 2016-11-14")
                os.Exit(0)
        }

        if *ip == "" {
                fmt.Println("No ntp server ip specified")
                os.Exit(1)
        }

        t, err := ntpclient.GetNetworkTime(*ip, *port)
        if err != nil {
                fmt.Println("Get ntp time error")
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Println(int(t.Sub(time.Now()).Seconds()))
}