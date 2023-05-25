package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

const ExitCodeBadArg = 2
const ExitCodeBadHost = 3
const ExitCodeBadPort = 4

var (
	versionFlag = flag.Bool("version", false, "Show version")
	tcpFlag = flag.Bool("tcp", false, "Send all packets to TCP")
	udpFlag = flag.Bool("udp", false, "Send all packets to UDP")
	delayFlag = flag.Int("delay", 5, "Delay between each packet")

	version = "develop"
	appName = "OpenGoKnocking"
)

type packet struct {
	Port int
	Protocol string
}

func showUsage() {
	showVersion()
	fmt.Fprintf(flag.CommandLine.Output(), `
Usage of OpenGoKnocking:

  Knock to port with many protocol:
    ./open-go-knocking <host> <port[:protocol]> <port[:protocol]> <port[:protocol]> ...
    ./open-go-knocking 127.0.0.1 1337:tcp 1338:udp 1339:tcp

  Knock to port only tcp:
    ./open-go-knocking --tcp <host> <port> <port> <port> ...

  Knock to port only udp:
    ./open-go-knocking --udp <host> <port> <port> <port> ...

  Show this help
    ./open-go-knocking -h|--help

  Show OpenGoKnocking version
    ./open-go-knocking --version

Options:
`)
	flag.PrintDefaults()
}

func showVersion() {
	fmt.Fprintf(flag.CommandLine.Output(), "%s (%s)\n", appName, version)
}

func exitWithError(errorCode int, errorMessage string, showHelp bool) {
	fmt.Fprintf(flag.CommandLine.Output(), "ERROR: %s\n\n", errorMessage)

	if showHelp {
		showUsage()
	}

	os.Exit(errorCode)
}

func getHostAndVerify() (string) {
	host := flag.Arg(0)
	if !govalidator.IsHost(host) {
		exitWithError(ExitCodeBadHost, "Host is not a valid host (DNS name or IP v4/6) !", true)
	}

	return host
}

func getPacketsAndVerify() ([]packet) {
	var packets []packet
	var defaultProtocol string
	if *udpFlag {
		defaultProtocol = "udp"
	} else if *tcpFlag {
		defaultProtocol = "tcp"
	}

	arguments := flag.CommandLine.Args()
	max := len(arguments)

	for iterator := 1; iterator < max; iterator++ {
		packetDetail := strings.Split(arguments[iterator], ":");

		if len(packetDetail) == 1 && defaultProtocol == "" {
			exitWithError(ExitCodeBadPort, "each packet must have protocol or you must defined a default protocol !", true)
		}

		if len(packetDetail) == 2 && packetDetail[1] != "tcp" && packetDetail[1] != "udp" {
			exitWithError(ExitCodeBadPort, "protocol must be udp or tcp !", true)
		}

		port, err := strconv.Atoi(packetDetail[0])
		if err != nil {
			exitWithError(ExitCodeBadPort, "port must be an integer !", true)
		}
		
		if port < 1 || port > 65535 {
			exitWithError(ExitCodeBadPort, "port must be an integer between 1 and 65535 !", true)
        }

		if len(packetDetail) == 2 {
			packets = append(packets, packet{Port: port, Protocol: packetDetail[1]})
		} else {
			packets = append(packets, packet{Port: port, Protocol: defaultProtocol})
		}
	}

	return packets
}

func knock(host string, packet packet) {
	destination := net.JoinHostPort(host, strconv.Itoa(packet.Port))
	
	conn, err := net.DialTimeout(packet.Protocol, destination, 5 * time.Millisecond)
	if err != nil {
		// Expected error
	} else {
		defer conn.Close()
	}
	
    time.Sleep(time.Duration(*delayFlag)*time.Millisecond)
}


func main() {
	flag.Usage = showUsage
	flag.Parse()

	if *versionFlag {
		showVersion()
		os.Exit(0)
	}

	if flag.CommandLine.NArg() < 2 {
		exitWithError(1, "Arguments missing !", true)
	}

	host := getHostAndVerify()
	packets := getPacketsAndVerify()

	for _, packet := range packets {
		knock(host, packet)
	}

	os.Exit(0)
}
