package helper

import (
	"project/mathutils"
	"project/stringutils"
)

func Add(a, b int) int {
	return mathutils.Add(a, b)
}

func Length(s string) int {
	return stringutils.Length(s)
}
