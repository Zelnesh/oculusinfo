package portscanner

import (

	"net"
	"time"
	"context"
	"strings"
	"fmt"
	"sync"
	"strconv"
)





func scan(ip, port string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	address := net.JoinHostPort(ip,port)
	var dial net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := dial.DialContext(ctx, "tcp", address)
	if err != nil {
		results <- fmt.Sprintf("%s:%s -> CLOSE:%v", ip,port,err)
		return
	}
	defer conn.Close()


	results <- fmt.Sprintf("%s:%s -> OPEN", ip,port)
}


func ScanPort (ip, port string) <-chan string {

	var wg sync.WaitGroup
	results := make(chan string)
	if strings.Contains(port, ",") {
		ports := strings.Split(port, ",")

		go func(){
			for _, port := range ports {
			    wg.Add(1)
			    go scan(ip,strings.TrimSpace(port), results, &wg)
			}
			wg.Wait()
			close(results)
		}()
	}

	if strings.Contains(port, "-"){
		ports := strings.Split(port, "-")
		startPort, _ := strconv.Atoi(ports[0])
		endPort, _ := strconv.Atoi(ports[1])

		go func(){
			for port := startPort; port <= endPort; port++{
				wg.Add(1)
				go scan(ip, strconv.Itoa(port), results, &wg)
			}
			wg.Wait()
			close(results)

		}()
	}

	return results
}
