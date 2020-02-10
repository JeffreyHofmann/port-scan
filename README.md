# port-scan
Go Port Scanner

Usage of ./port-scan:
  -ep int
    	start port (default -1)
  -host string
    	REQUIRED - host to scan
  -p-
    	scan all ports
  -sp int
    	start port (default -1)
  -t int
    	number of threads to use (default 50)

`./port-scan -host scanme.nmap.org -p- -t 10`