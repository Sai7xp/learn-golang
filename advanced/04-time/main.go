/*
* Created on 23 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
const (

	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"

)
*/
func main() {
	fmt.Println("Time in GO lang")

	now := time.Now()

	fmt.Println("Current Date Time : ", now)
	fmt.Println("Current Date Time RFC3339 Format : ", now.Format(time.RFC3339))
	fmt.Println("ANSIC Formatted date time : ", now.Format(time.ANSIC))
	fmt.Println("UnixDate Formatted date time : ", now.Format(time.UnixDate))
	fmt.Println("Custom Formatted date time : ", now.Format("01/02/2006, Mon and time is -> 15:04 PM"))
	/// yes 01/02/2006 these all values are fixed and that's how we format the date

	fmt.Println()

	createdDateTime := time.Date(2001, time.March, 23, 10, 12, 59, 100000, time.Local)

	fmt.Println("createdDateTime ", createdDateTime)
	fmt.Println("createdDateTime ", createdDateTime.Format("01/02/2006, Mon and time is -> 15:04 PM"))
	milli := time.Now().UnixMilli()
	fmt.Println(time.UnixMilli(milli).Format("02/01/2006, Mon -> 15:04:05 PM"))

	totalCores := runtime.NumCPU()
	fmt.Printf("This machine has %d CPU cores. \n", totalCores)
	runtime.GOMAXPROCS(totalCores)

}
