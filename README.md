# port-scan
Go Port Scanner

Usage of ./port-scan: \
&nbsp;&nbsp;&nbsp;&nbsp;-ep int \
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;start port (default -1) \
&nbsp;&nbsp;&nbsp;&nbsp;-host string \
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;REQUIRED - host to scan \
&nbsp;&nbsp;&nbsp;&nbsp;-p- \
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;scan all ports \
&nbsp;&nbsp;&nbsp;&nbsp;-sp int \
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;start port (default -1) \
&nbsp;&nbsp;&nbsp;&nbsp;-t int \
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;number of threads to use (default 50)

`./port-scan -host scanme.nmap.org -p- -t 10`
