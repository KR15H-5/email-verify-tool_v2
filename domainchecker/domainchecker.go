package domainchecker

import (
	"log"
	"net"
	"strings"
)

type CheckResponse struct {
	Domain      string // The domain that was checked
	HasMX       bool   // Whether the domain has MX records
	HasSPF      bool   // Whether the domain has SPF records
	SPFRecord   string // The actual SPF record content
	HasDMARC    bool   // Whether the domain has DMARC records
	DMARCRecord string // The actual DMARC record content
	Error       string // To capture any errors that occurred during the check
}

func CheckDomain(domain string) CheckResponse {
	var response CheckResponse
	response.Domain = domain

	// Check for MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for %v: %v", domain, err)
		response.Error = "Failed to look up MX records"
	} else if len(mxRecords) > 0 {
		response.HasMX = true
	}

	// Check for SPF records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for %v: %v", domain, err)
		if response.Error == "" {
			response.Error = "Failed to look up TXT records"
		} else {
			response.Error += "; Failed to look up TXT records"
		}
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			response.HasSPF = true
			response.SPFRecord = record
			break
		}
	}

	// Check for DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for %v: %v", domain, err)
		if response.Error == "" {
			response.Error = "Failed to look up DMARC records"
		} else {
			response.Error += "; Failed to look up DMARC records"
		}
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			response.HasDMARC = true
			response.DMARCRecord = record
			break
		}
	}

	return response
}
