package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func main() {
	var (
		daemon  = flag.String("d", "", "FRR daemon to run the command in")
		command = flag.String("c", "", "command to run")
	)
	flag.Parse()

	if *daemon == "" {
		fmt.Println("Error: daemon parameter required")
		return
	}

	socketPath, err := lookupSocketPath(*daemon)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	outb, err := runCmd(socketPath, *command)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Print(string(outb))
}

func lookupSocketPath(daemon string) (string, error) {
	switch daemon {
	case
		"babeld",
		"bfdd",
		"bgpd",
		"eigrpd",
		"fabricd",
		"isisd",
		"ldpd",
		"nhrpd",
		"ospf6d",
		"ospfd",
		"pbrd",
		"pimd",
		"ripd",
		"ripngd",
		"sharpd",
		"staticd",
		"vrrpd",
		"zebra":
		return fmt.Sprintf("/var/run/frr/%s.vty", daemon), nil
	}
	return "", fmt.Errorf("unknown daemon %s", daemon)
}

func runCmd(socketPath string, cmd string) ([]byte, error) {
	socket, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}
	defer socket.Close()

	cmd = cmd + "\x00"
	_, err = socket.Write([]byte(cmd))
	if err != nil {
		return nil, err
	}

	output, err := bufio.NewReader(socket).ReadBytes('\x00')
	if err != nil {
		return nil, err
	}

	return output[:len(output)-1], nil
}
