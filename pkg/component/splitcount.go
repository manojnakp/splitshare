package component

import "time"

type SplitCount struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	By           Name      `json:"by"`
	On           time.Time `json:"on"`
	Participants []Name    `json:"participants"`
}
