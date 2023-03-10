package model

import "fmt"

type Content struct {
	Title    string
	Date     string
	PosScore int
	NegScore int
}

func NewContent(title, date string, posScore, negScore int) *Content {
	return &Content{
		Title:    title,
		Date:     date,
		PosScore: posScore,
		NegScore: negScore,
	}
}

func (c *Content) ToSlice() []string {
	contentSlice := make([]string, 0)
	contentSlice = append(contentSlice, c.Date, c.Title, fmt.Sprintf("%d", c.PosScore), fmt.Sprintf("%d", c.NegScore))
	return contentSlice
}
