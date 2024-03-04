package main

import (
	"github.com/maxguuse/100GoMistakePractice/13-UtilsPackage/stringset"
	"github.com/maxguuse/100GoMistakePractice/13-UtilsPackage/utils"
)

func main() {
	s := utils.NewStringSet()
	_ = utils.SortStringSet(s) // Name of package doesn't tell us anything

	s2 := stringset.New() // Better
	_ = s2.Sort()
}
