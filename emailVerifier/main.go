package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// checkDomain function checks various DNS records for a given domain and prints the results.
func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool // Flags to track presence of MX, SPF, and DMARC records
	var sprRecord, dmarcRecord string // Store SPF and DMARC records if found

	// Lookup MX records for the domain
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error fetching MX record: %v \n", err)
	}
	if len(mxRecord) > 0 {
		hasMx = true // Set hasMx flag if MX records are found
	}

	// Lookup TXT records for the domain
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("TXT Record Error: %v \n", err)
	}

	// Check for SPF record in TXT records
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true       // Set hasSPF flag if SPF record is found
			sprRecord = record  // Store SPF record for display
			break
		}
	}

	// Lookup DMARC records for the domain
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error in DMARC record check: %v\n", err)
	}

	// Check for DMARC record in DMARC records
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true        // Set hasDMARC flag if DMARC record is found
			dmarcRecord = record   // Store DMARC record for display
			break
		}
	}

	// Print formatted output for the domain's DNS records
	fmt.Printf("Domain: %s\nHas MX Record: %v\nHas SPF Record: %v\nSPF Record: %s\nHas DMARC Record: %v\nDMARC Record: %s\n\n",
		domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	// Read domains from stdin and check their DNS records
	for scanner.Scan() {
		checkDomain(scanner.Text()) // Invoke checkDomain function for each domain entered
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
}
