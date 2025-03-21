package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindAllRecipes$
func TestFindAllRecipes(t *testing.T) {
	for _, tc := range []struct {
		recipes        []string
		ingredients    [][]string
		supplies       []string
		createdRecipes []string
	}{
		{recipes: []string{"bread"}, ingredients: [][]string{{"yeast", "flour"}}, supplies: []string{"yeast", "flour", "corn"}, createdRecipes: []string{"bread"}},
		{recipes: []string{"bread", "sandwich"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}}, supplies: []string{"yeast", "flour", "meat"}, createdRecipes: []string{"bread", "sandwich"}},
		// {recipes: []string{"bread", "sandwich", "burger"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, supplies: []string{"yeast", "flour", "meat"}, createdRecipes: []string{"bread", "sandwich", "burger"}},
		// {recipes: []string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}, ingredients: [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}}, supplies: []string{"f", "hveml", "cpivl", "d"}, createdRecipes: []string{"fzjnm", "q", "ju"}},
		{recipes: []string{"bread", "sandwich"}, ingredients: [][]string{{"bread", "flour"}, {"bread", "flour"}}, supplies: []string{"yeast", "flour", "meat"}, createdRecipes: []string{}},
		// {recipes: []string{"qxyj", "vawos", "nkov", "bec", "qiabz"}, ingredients: [][]string{{"mxf"}, {"iy", "qxyj", "nkov", "qiabz", "bec"}, {"nw", "xutnl", "e"}, {"eep", "km", "nw", "xutnl", "e", "iy", "vawos", "qxyj", "qiabz"}, {"nyhyc"}}, supplies: []string{"nw", "eep", "iy", "e", "xutnl", "km"}, createdRecipes: []string{"nkov"}},
		{recipes: []string{"bread", "sandwich"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "flour"}}, supplies: []string{"yeast", "flour", "meat"}, createdRecipes: []string{"bread", "sandwich"}},
	} {
		createdRecipes := findAllRecipes(tc.recipes, tc.ingredients, tc.supplies)
		assert.Equal(t, tc.createdRecipes, createdRecipes)
	}
}
