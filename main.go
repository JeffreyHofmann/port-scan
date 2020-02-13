package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- -1
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	var (
		host      string
		startPort int
		endPort   int
		threads   int
		allPorts  bool
	)

	flag.StringVar(&host, "host", "", "REQUIRED - host to scan")
	flag.IntVar(&startPort, "sp", -1, "start port")
	flag.IntVar(&endPort, "ep", -1, "start port")
	flag.IntVar(&threads, "t", 50, "number of threads to use")
	flag.BoolVar(&allPorts, "p-", false, "scan all ports")
	flag.Parse()

	if allPorts {
		startPort = 1
		endPort = 65535
	}

	if startPort > endPort || startPort < 1 || endPort > 65535 {
		startPort = -1
	}

	if host == "" || endPort == -1 || startPort == -1 {
		flag.Usage()
		os.Exit(1)
	}

	ports := make(chan int, threads)
	defer close(ports)
	results := make(chan int)
	defer close(results)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, results)
	}

	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
	}()

	stdout := bufio.NewWriter(os.Stdout)
	for i := startPort; i <= endPort; i++ {
		port := <-results
		if port != -1 {
			openports = append(openports, port)
		}
		fmt.Fprintf(stdout, "%s - %d/%d ports scanned\r",
			host, 1+i-startPort, 1+endPort-startPort)
		stdout.Flush()
	}

	sort.Ints(openports)
	fmt.Printf("\n%d/%d ports open\n", len(openports), 1+endPort-startPort)
	for _, port := range openports {
		fmt.Printf("port %d open\n", port)
	}
}
