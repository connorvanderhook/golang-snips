package main

import "fmt"

func translateDirection(from, to string) string {
	return fmt.Sprintf("%s-%s", from, to)
}
