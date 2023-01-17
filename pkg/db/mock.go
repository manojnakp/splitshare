package db

import (
	"github.com/manojnakp/splitshare/pkg/component"
	"time"
)

var Expenses Collection = &ECollection{{
	Id:     "CVYXKEJZWJE7PJMH4KDAWMGBD4",
	Name:   "Lunch",
	Amount: 300,
	By:     "Julia",
	When:   time.Now(),
	Among: []float64{
		120,
		60,
		60,
		60,
	},
}}

var SplitCounts Collection = &SCollection{{
	Id:    "kb9t18wbboczf8xecdiy",
	Title: "City Trip",
	Desc:  "Sample splitcount sharing expenses of a city trip among friends",
	By:    "Alex",
	On:    time.Now(),
	Participants: []component.Name{
		"Alex",
		"Brian",
		"Julia",
		"Thomas",
	},
}}
