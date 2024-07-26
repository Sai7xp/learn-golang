/*
* Created on 02 April 2024
* @author Sai Sumanth
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Blood Report Type
type BloodReport struct {
	Content string
}

/// now let's suppose we have one more report Type called 'EcgReport'. Now the problem is how can we use
// that inside GlobalExport Function

// Ecg Report Type
type EcgReport struct {
	Content []string
}

type UniversalReport interface {
	export()
}

func (br *BloodReport) export() {
	bytes, err := json.Marshal(br)

	if err == nil {
		// export the report to a file
		os.WriteFile("csv_report.json", bytes, 0644)
	}
}

func (br *EcgReport) export() {
	bytes, err := json.Marshal(br)

	if err == nil {
		// export the report to a file
		os.WriteFile("pdf_report.json", bytes, 0644)
	}
}

func OpenClosed() {
	fmt.Println("\n2️⃣ Open Closed Principle - Go Lang SOLID Principles")
	bloodReport := BloodReport{Content: "Blood Report Content"}
	ecgReport := EcgReport{Content: []string{"ECG 1", "ECG 2"}}

	// export reports
	globalExport(&bloodReport)
	globalExport(&ecgReport)

	ecgReport.exportToCSV()

}

func globalExport(ur UniversalReport) {
	ur.export()
}

// New method specific to EcgReport
func (er *EcgReport) exportToCSV() {
	bytes, err := json.Marshal(er)
	if err == nil {
		// export the report to a CSV file
		os.WriteFile("ecg_report.csv", bytes, 0644)
	}
}
