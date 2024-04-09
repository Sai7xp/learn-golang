/*
* Created on 2 April 2024
* @author Sai Sumanth
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/*
	❌ Violating the Single Responsibility Principle
*/

// report with content
type Report struct {
	Content string
}

// report file schema with generated time info, File will be generated as per this schema
type ReportFileSchema struct {
	R             Report
	GeneratedTime time.Time
}

// Creates Report and Generates the Report
//
// ❌ This function is responsible for two Actions Generating Report and Exporting Report, which
// violates the Single Responsibility Principle
func (r *Report) generateAndExportReport() {
	fmt.Println("Creating Report File Format with current time")
	/// Generate Report
	rSchema := ReportFileSchema{R: *r, GeneratedTime: time.Now()}

	bytes, _ := json.Marshal(rSchema)
	// export the report to a file
	os.WriteFile("report.json", bytes, 0644)

}

func main() {
	fmt.Println("Single Responsibility Principle - Go Lang")

	rep := Report{Content: "This report is generated by computer."}
	rep.generateAndExportReport()

	/////
	rSchema := rep.generateReport()
	rSchema.exportReportFile()
}

/*
	✅ Implementing Single Responsibility Principle

	This separation ensures that each method has a single responsibility
	and is less likely to change for multiple reasons.
*/

// function responsible only for Generating the report
func (r *Report) generateReport() ReportFileSchema {
	fmt.Println("Creating Report File Format with current time")

	/// Generate Report
	rSchema := ReportFileSchema{R: *r, GeneratedTime: time.Now()}
	return rSchema
}

// function responsible only for Exporting the report
func (rSchema *ReportFileSchema) exportReportFile() {
	bytes, err := json.Marshal(rSchema)

	if err == nil {
		// export the report to a file
		os.WriteFile("report_pro.json", bytes, 0644)
	}
}
