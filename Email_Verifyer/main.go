package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the domain address you want to verify:")

	for scanner.Scan() {
		checkDomain(scanner.Text())
		// fmt.Println("\nEnter another domain or press Ctrl+C to exit:")
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check MX record
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v", domain, err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check SPF record
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v", domain, err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// Check DMARC record
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for domain %s: %v", domain, err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// Output Results
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf(" - Has MX record: %t\n", hasMX)
	fmt.Printf(" - Has SPF record: %t\n", hasSPF)
	if hasSPF {
		fmt.Printf("   SPF Record: %s\n", spfRecord)
	}
	fmt.Printf(" - Has DMARC record: %t\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("   DMARC Record: %s\n", dmarcRecord)
	}
}
