package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	reportDateNumeric := strings.ReplaceAll(reportDate, "-", "")
	
	reportIDStr := strconv.Itoa(user.ID) + reportDateNumeric
	
	reportID, _ := strconv.Atoi(reportIDStr)

	return Report{
		User:     user,
		ReportID: reportID,
		Date:     reportDate,
	}
}

func PrintReport(report Report) {
	fmt.Printf("Name: %s, Age: %d, ReportID: %d, Date: %s\n", report.Name, report.Age, report.ReportID, report.Date)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0)
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}
	return reports
}

func main() {
	dani := User{
		ID:    777,
		Name:  "Danial",
		Email: "doniponi@gmail.com",
		Age:   18,
	}
	bob := User{
		ID:    100,
		Name:  "Bobik",
		Email: "bobikshobik@gmail.com",
		Age:   10,
	}
	users := []User{dani, bob}
	reportDate := time.Now().Format("2006-01-02")
	reports := GenerateUserReports(users, reportDate)

	for _, report := range reports {
		PrintReport(report)
	}
}
