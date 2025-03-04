package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func main() {
	examples := []string{
		"recommendedAscentSpeed2",
		"HTTPServer",
		"IPAddress",
		"someVariable42",
		"convertToSnakeCase",
	}

	/*
		recommendedAscentSpeed2 -> recommended_ascent_speed_2
		HTTPServer -> http_server
		IPAddress -> ip_address
		someVariable42 -> some_variable_42
		convertToSnakeCase -> convert_to_snake_case
	*/
	for _, example := range examples {
		fmt.Printf("%s -> %s\n", example, strcase.ToSnake(example))
	}
}
