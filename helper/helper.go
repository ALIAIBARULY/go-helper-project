package helper

import (
	"github.com/ALIAIBARULY/go-helper-project/mathutils"
	"github.com/ALIAIBARULY/go-helper-project/stringutils"
)

func Add(a, b int) int {
	return mathutils.Add(a, b)
}

func Length(s string) int {
	return stringutils.Length(s)
}
