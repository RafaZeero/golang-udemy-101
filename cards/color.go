package main

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}
