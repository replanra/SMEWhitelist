package config

import (
	"html/template"
	"math"
	"time"

	"github.com/dustin/go-humanize"
)

var TPL *template.Template
var TPLGO *template.Template

func init() {
	TPL = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
	//	TPLGO = template.Must(template.New("").Funcs(fm).ParseGlob("templates/gohtml/*.gohtml"))
}



//tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"money":    money,
}

func monthDayYear(t time.Time) string {
	//return t.Format("01-02-2006")
	return t.Format("02 Jan 2006")

}

func money(f float32) string {
	return humanize.Commaf(float64(math.Round(float64(f)*100) / 100))
}
