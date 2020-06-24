package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func pingAus(ip string, c chan string, wg *sync.WaitGroup) {
	// ping ausführen
	defer wg.Done()                                       // sicherstellen das wenn der durchlauf durch ist, das aus der Waitgroup der job rausgenommen wird
	out, err := exec.Command("ping", "-c 4", ip).Output() // Ausführen des Commandos
	if err != nil {                                       // Fehler Abfangen
		log.Println("Ohh ein Fehler")
	}
	c <- fmt.Sprintf("%s", out) // Rückgabe der Ausgabe als String ni den Channel c
}

func main() {
	var wg sync.WaitGroup // Waitgoup zum Syncen definieren

	ips := []string{"127.0.0.1", "8.8.8.8", "1.1.1.1", "192.168.1.1", "192.168.1.13"} // IPs

	c := make(chan string) // Channel definieren

	for _, a := range ips { //Schleife der IP Adressen durchlaufen und jeweils
		wg.Add(1) // den Counter in der WaitGroup erhöhen in
		go pingAus(a, c, &wg)
	}

	go func() { // in einer GO Routine wird die WaitGroup auf warten und Channel close geschickt
		wg.Wait()
		close(c)
	}()

	for out := range c { // Ergebnis aus dem Channel rausholen
		fmt.Println("##################")
		fmt.Printf("%s", out)
	}
}
