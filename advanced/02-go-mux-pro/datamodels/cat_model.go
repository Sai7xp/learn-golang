package datamodels

import (
	"fmt"
	"time"
)

type Cat struct {
	Name         string    `json:"catName"`
	AgeInMonths  uint      `json:"ageInMonths"`
	WeightInKgs  float64   `json:"weight"`
	Color        string    `json:"color"`
	IsMale       bool      `json:"-"` // this field will be hidden in json conversions
	DateOfBirth  time.Time `json:"dob"`
	Vaccinations []string  `json:"vaccinations"`
	Owner        *Owner    `json:"catOwner"`
}

func (c *Cat) IsVaccinated() bool {
	return len(c.Vaccinations) > 0
}

// PrintInfo prints the information about the cat
func (c *Cat) PrintInfo() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Age in months: %d\n", c.AgeInMonths)
	fmt.Printf("Weight in kgs: %.2f\n", c.WeightInKgs)
	fmt.Printf("Color: %s\n", c.Color)
	fmt.Printf("Date of Birth: %s\n", c.DateOfBirth.Format("2006-01-02"))
	fmt.Println("Vaccinations:")
	for _, v := range c.Vaccinations {
		fmt.Println("-", v)
	}
	fmt.Printf("Is Male: %t\n", c.IsMale)
	fmt.Printf("Owner Is: %s\n", c.Owner.Name)
}

func (c *Cat) ToString() string {
	res := fmt.Sprintf("Name %s, Age: %d, Date of Birth: %s", c.Name, c.AgeInMonths, c.DateOfBirth)
	return res
}

type Owner struct {
	Name   string `json:"catOwner"`
	IsMale bool   `json:"isCatOwnerMale"`
}
