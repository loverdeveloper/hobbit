package pdco_go

import (
	"encoding/json"
	"fmt"
)

type Akbar struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Grade  string `json:"grade"`
}

func (a *Akbar) getAkbar() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Akbar) setName(name string)  {
	a.Name = name
}

func (a *Akbar) setFamily(family string)  {
	a.Family = family
}

func (a *Akbar) setGrade(grade string)  {
	a.Grade = grade
}

func akbarFactory()  {
	var akbar Akbar
	akbar.setName("Akbar")
	akbar.setFamily("Ghaeini")
	akbar.setGrade("Grade Zahra Khodayary")
	akbarIns, _ := akbar.getAkbar()
	fmt.Println(string(akbarIns))
}

func main() {
	akbarFactory()
}