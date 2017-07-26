package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Libelle              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y1 := year{
		Libelle: "2016",
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming", "5"},
				course{"CSCI-130", "Introduction to Wordpress", "4"},
				course{"CSCI-140", "Mobile Apps Using AndroidStudio", "5"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming", "5"},
				course{"CSCI-191", "Advanced Mobile Apps", "4"},
			},
		},
	}

	y2 := year{
		Libelle: "2017",
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming", "5"},
				course{"CSCI-130", "Introduction to Wordpress", "4"},
				course{"CSCI-140", "Mobile Apps Using AndroidStudio", "5"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming", "5"},
				course{"CSCI-191", "Advanced Mobile Apps", "4"},
			},
		},
	}

	years := []year{y1, y2}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}
