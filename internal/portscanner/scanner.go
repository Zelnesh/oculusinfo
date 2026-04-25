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


func portArgInputCharValidation(s string) bool {
	for _, ch := range s{
		if (ch >= '0' && ch <= '9') || ch == '-' || ch == ',' {
			continue
		}
		return false
	}
	return true
}


func ScanPort (ip, port string) <-chan string {

	var wg sync.WaitGroup
	results := make(chan string)
	if !portArgInputCharValidation(port){
		go func(){
			results <- "Wrong port argument. Please check documentation using --help"
			close(results)
		}()
		return results
	} else if strings.Contains(port, "-") && strings.Contains(port, ","){
		go func(){
			results <- "Wrong port argument. Please check documentation using --help"
			close(results)
		}()
		return results

	}else if strings.Contains(port, ",") {
		ports := strings.Split(port, ",")
		for _, p := range ports {
			if p == "" {
				go func(){
					results <- "Worng port argument. Please check ocumentation using --help"
					close(results)
				}()
				return results
			}
		}

		go func() {
			for _, port := range ports {
				wg.Add(1)
				go scan(ip, strings.TrimSpace(port), results, &wg)
			}
			wg.Wait()
			close(results)
		}()

	} else if strings.Contains(port, "-") {
		ports := strings.Split(port, "-")
		if len(ports) != 2 {
			go func(){
				results <- "Worng port range argument. Please check documentation using --help"
				close(results)
			}()
			return results
		}
		startPort, _ := strconv.Atoi(ports[0])
		endPort, _ := strconv.Atoi(ports[1])

		go func() {
			for port := startPort; port <= endPort; port++ {
				wg.Add(1)
				go scan(ip, strconv.Itoa(port), results, &wg)
			}
			wg.Wait()
			close(results)
		}()
	} else {
		portCheck, err := strconv.Atoi(port)
		if err != nil {
			go func(){
				results <- "Port should be an integer number..."
				close(results)
			}()
			return results
		}
		if portCheck < 1 || portCheck > 65535 {
			go func(){
				results <- "Port out of range, range could be 1-65535..."
				close(results)
			}()
			return results
		}
		go func(){
			wg.Add(1)
			go scan(ip, strconv.Itoa(portCheck), results, &wg)
			wg.Wait()
			close(results)

		}()
	}


	return results
}
