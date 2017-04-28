package main

// Written by: Steven Shatz on 4/27/2017

// Project Euler - Problem 19
//  How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
//  Given: 1 Jan 1900 was a Monday
//  Expected Answer = 171

import (
	"fmt"
	"time"
)

// Notes on "time.go":
// time.Month() returns the full name of the current month
// time.Day() returns the current day of the month
// time.Year() returns the current year
// time.Month(m) returns the full name of the specified month (1-12) as a string (eg: time.Month(1) returns "January")
// time.Local is the local timezone
// time.Date(yyyy, mm, dd, hh, mm, ss, nn, time.Local) returns a "time.go" Date structure.  (Note: nn=nanoseconds)
// time.Weekday() returns the full name of the corresponding day of the week as a string (eg: "Sunday")

func main() {
	fromDate := setDate(1, 1, 1901)
	toDate := setDate(12, 31, 2000)

	fmt.Printf("\nFrom Date is %s %d, %d \n", fromDate.Month(), fromDate.Day(), fromDate.Year())
	fmt.Printf("To Date   is %s %d, %d \n", toDate.Month(), toDate.Day(), toDate.Year())

	fmt.Printf("There were %d Sundays on the first of a month in the Twentieth Century \n\n", countSundaysInRange(fromDate, toDate))
}

// setDate converts a specified month, day, year into a "time.go" Date structure.
func setDate(m int, day int, year int) time.Time {
	return time.Date(year, time.Month(m), day, 0, 0, 0, 0, time.Local)
}

// countSundaysInRange counts the number of Sundays between the from and to Dates, inclusive.
func countSundaysInRange(fromDate time.Time, toDate time.Time) int {
	var totalSundays int
	aDay := time.Duration(24 * time.Hour)
	testDate := fromDate
	for !testDate.After(toDate) {
		if testDate.Day() == 1 && testDate.Weekday() == time.Sunday {
			totalSundays++
		}
		testDate = testDate.Add(aDay)
	}
	return totalSundays
}
