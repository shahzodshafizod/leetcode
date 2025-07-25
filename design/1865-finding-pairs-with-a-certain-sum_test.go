package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestFindSumPairs$
func TestFindSumPairs(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][][]int
		output   []any
	}{
		{
			commands: []string{"FindSumPairs", "count", "add", "count", "count", "add", "add", "count"},
			values:   [][][]int{{{1, 1, 2, 2, 2, 3}, {1, 4, 5, 2, 5, 4}}, {{7}}, {{3, 2}}, {{8}}, {{4}}, {{0, 1}}, {{1, 1}}, {{7}}},
			output:   []any{nil, 8, nil, 2, 1, nil, nil, 11},
		},
		// {
		// 	commands: []string{"FindSumPairs","add","add","count","add","add","add","add","add","add","add","add","count","count","add","add","add","add","add","add","add","add","add","count","add","add","count","add","add","add","add","count","count","add","add","add","count","add","count","add","add","add","count","add","count","add","add","add","add","add","count","add","add","add","add","count","add","count","add","count","add","add","add","add","add","add","add","add","count","add","add","add","add","add","count","add","add","count","add","add","add","add","add","add","add","count","add","add","count","add","count","add","add","add","count","add","add","add","add","add","add","add","add","add","add","add","add","add","count","add","add","add","add","add","add","add","count","add","count","add","add","add","add","count","count","add","add","add","add","add","add","add","add","add","add","count","add","add","add","add","count","add","count","add","add","add","add","add","add","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count"},
		// 	values: [][][]int{{{9,70,14,9,76},{26,26,58,23,74,68,68,78,58,26}},{{6,10}},{{5,6}},{{32}},{{3,55}},{{9,32}},{{9,16}},{{1,48}},{{1,4}},{{0,52}},{{8,20}},{{9,4}},{{88}},{{154}},{{9,4501}},{{2,20}},{{2,4501}},{{8,4501}},{{5,4505}},{{4,4}},{{5,4241}},{{4,4501}},{{4,4241}},{{4588}},{{0,4501}},{{4,847}},{{8896}},{{1,9589}},{{5,847}},{{0,5088}},{{3,9589}},{{4649}},{{92}},{{6,4501}},{{6,5088}},{{7,9589}},{{4649}},{{7,9671}},{{9681}},{{6,9671}},{{8,5088}},{{2,5088}},{{19347}},{{9,5088}},{{9681}},{{9,9671}},{{2,9671}},{{2,7295}},{{9,7295}},{{8,9671}},{{19347}},{{3,9671}},{{0,9671}},{{5,9671}},{{1,9671}},{{9676}},{{4,9671}},{{19408}},{{4,7295}},{{26642}},{{1,7295}},{{5,7295}},{{5,3113}},{{1,3113}},{{0,7295}},{{3,7295}},{{8,7295}},{{6,7295}},{{29816}},{{7,7295}},{{7,3113}},{{7,2215}},{{6,3113}},{{8,3113}},{{29755}},{{3,3113}},{{0,5328}},{{32037}},{{4,3113}},{{9,5328}},{{2,3113}},{{2,2215}},{{4,2215}},{{4,6664}},{{3,2215}},{{38639}},{{4,3574}},{{8,2215}},{{29755}},{{6,2215}},{{29816}},{{4,4791}},{{4,8635}},{{4,2734}},{{29760}},{{1,2215}},{{5,2215}},{{4,9817}},{{4,4323}},{{4,7213}},{{4,6589}},{{4,1480}},{{4,9856}},{{4,8968}},{{4,8194}},{{4,4939}},{{4,7355}},{{4,8997}},{{31975}},{{4,6130}},{{4,9464}},{{4,7884}},{{4,9954}},{{4,2439}},{{4,4887}},{{4,7171}},{{184028}},{{4,1721}},{{32031}},{{4,9642}},{{4,3381}},{{4,2846}},{{4,3498}},{{205183}},{{205183}},{{4,6018}},{{4,762}},{{4,8810}},{{4,510}},{{4,6949}},{{4,949}},{{4,9201}},{{4,4371}},{{4,3002}},{{4,4107}},{{32037}},{{4,1863}},{{4,8839}},{{4,8822}},{{4,9054}},{{32031}},{{4,2630}},{{31970}},{{4,7110}},{{4,7373}},{{4,271}},{{4,5630}},{{4,3345}},{{4,1512}},{{306244}},{{306249}},{{31970}},{{32031}},{{306305}},{{306311}},{{31970}},{{31970}},{{306311}},{{306249}},{{306305}},{{306244}},{{306249}},{{306244}},{{32031}},{{306249}},{{31970}},{{32031}},{{31975}},{{31975}},{{306244}},{{31975}},{{306311}},{{32031}},{{306311}},{{306244}},{{306249}},{{31975}},{{306311}},{{32031}},{{306244}},{{306244}},{{306305}},{{306311}},{{32031}},{{31970}},{{31970}},{{306305}},{{306244}},{{31970}},{{306311}}},
		// 	output: []any{nil,nil,nil,2,nil,nil,nil,nil,nil,nil,nil,nil,2,7,nil,nil,nil,nil,nil,nil,nil,nil,nil,6,nil,nil,1,nil,nil,nil,nil,3,2,nil,nil,nil,3,nil,6,nil,nil,nil,4,nil,8,nil,nil,nil,nil,nil,6,nil,nil,nil,nil,2,nil,8,nil,6,nil,nil,nil,nil,nil,nil,nil,nil,2,nil,nil,nil,nil,nil,8,nil,nil,2,nil,nil,nil,nil,nil,nil,nil,1,nil,nil,6,nil,2,nil,nil,nil,2,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,9,nil,nil,nil,nil,nil,nil,nil,2,nil,9,nil,nil,nil,nil,1,1,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,9,nil,nil,nil,nil,9,nil,18,nil,nil,nil,nil,nil,nil,2,1,18,9,1,1,18,18,1,1,1,2,1,2,9,1,18,9,9,9,2,9,1,9,1,2,1,9,1,9,2,2,1,1,9,18,18,1,2,18,1},
		// },
		// {
		// 	commands: []string{"FindSumPairs","add","add","count","add","add","add","add","add","add","count","add","add","add","add","add","add","count","add","add","add","count","add","add","count","count","add","add","add","add","add","add","add","add","add","count","count","add","add","add","add","add","count","add","add","add","add","add","add","add","add","add","add","add","add","add","count","add","add","add","count","add","add","add","add","add","add","count","add","add","add","count","add","add","add","add","add","add","add","add","add","add","add","add","add","add","count","add","add","add","add","add","count","add","count","add","add","add","count","count","add","add","add","add","add","add","add","add","add","add","add","add","count","add","add","add","add","add","add","count","count","add","add","add","count","count","add","add","add","add","count","add","add","add","add","count","count","add","count","add","add","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count","count"},
		// 	values: [][][]int{{{74},{90,90,84,84,90,90,90,90,90,90,90,84,90,90,84}},{{14,6}},{{11,6}},{{164}},{{11,2588}},{{14,2588}},{{14,9992}},{{11,9992}},{{3,6}},{{11,3042}},{{12744}},{{11,8319}},{{2,6}},{{11,7961}},{{11,2720}},{{11,3374}},{{11,768}},{{164}},{{11,8889}},{{11,1763}},{{11,5781}},{{164}},{{11,3286}},{{11,9934}},{{12744}},{{12744}},{{11,5072}},{{11,3606}},{{11,3980}},{{11,5474}},{{11,5663}},{{11,6814}},{{11,8376}},{{11,1026}},{{11,3780}},{{112372}},{{12744}},{{11,6776}},{{11,3986}},{{11,2778}},{{11,1930}},{{11,7070}},{{12744}},{{11,3409}},{{11,8546}},{{11,3497}},{{11,5176}},{{11,9478}},{{11,6823}},{{11,8350}},{{11,8298}},{{11,8616}},{{11,2676}},{{11,14}},{{11,1481}},{{11,4033}},{{164}},{{11,4277}},{{11,3325}},{{11,2108}},{{215019}},{{11,3832}},{{11,8610}},{{11,3471}},{{11,8321}},{{11,1175}},{{11,9045}},{{12744}},{{11,4424}},{{11,2862}},{{11,2318}},{{259077}},{{11,3236}},{{11,4950}},{{11,4117}},{{11,685}},{{11,4632}},{{11,6733}},{{11,7117}},{{11,9775}},{{11,8495}},{{11,3664}},{{11,2000}},{{11,9162}},{{11,728}},{{11,9544}},{{333915}},{{11,5461}},{{11,6541}},{{11,2460}},{{11,5823}},{{11,8438}},{{362638}},{{11,3720}},{{12744}},{{11,8906}},{{11,613}},{{11,5752}},{{164}},{{381629}},{{11,6398}},{{11,9719}},{{11,6133}},{{11,6778}},{{11,9822}},{{11,8134}},{{11,5158}},{{11,2998}},{{11,3740}},{{11,6078}},{{11,4394}},{{11,8020}},{{12744}},{{11,5307}},{{11,64}},{{11,299}},{{11,4646}},{{11,1967}},{{11,2863}},{{474147}},{{164}},{{11,8618}},{{11,3474}},{{11,7285}},{{164}},{{12744}},{{11,7687}},{{11,6589}},{{11,9010}},{{11,4772}},{{521582}},{{11,1334}},{{11,5078}},{{11,5432}},{{11,5418}},{{12744}},{{164}},{{11,2647}},{{541491}},{{11,5213}},{{11,1979}},{{548683}},{{12744}},{{12744}},{{548683}},{{164}},{{12744}},{{164}},{{12744}},{{548683}},{{12744}},{{12744}},{{548683}},{{548683}},{{548683}},{{164}},{{548683}},{{548683}},{{164}},{{12744}},{{164}},{{548683}},{{164}},{{164}},{{12744}},{{164}},{{164}},{{12744}},{{12744}},{{164}},{{548683}},{{164}},{{12744}},{{548683}},{{164}},{{548683}},{{12744}},{{12744}},{{164}},{{12744}},{{164}},{{164}},{{12744}},{{164}},{{164}},{{12744}},{{164}},{{164}},{{12744}},{{164}}},
		// 	output: []any{nil,nil,nil,13,nil,nil,nil,nil,nil,nil,1,nil,nil,nil,nil,nil,nil,13,nil,nil,nil,13,nil,nil,1,1,nil,nil,nil,nil,nil,nil,nil,nil,nil,1,1,nil,nil,nil,nil,nil,1,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,13,nil,nil,nil,1,nil,nil,nil,nil,nil,nil,1,nil,nil,nil,1,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,1,nil,nil,nil,nil,nil,1,nil,1,nil,nil,nil,13,1,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,1,nil,nil,nil,nil,nil,nil,1,13,nil,nil,nil,13,1,nil,nil,nil,nil,1,nil,nil,nil,nil,1,13,nil,1,nil,nil,1,1,1,1,13,1,13,1,1,1,1,1,1,1,13,1,1,13,1,13,1,13,13,1,13,13,1,1,13,1,13,1,1,13,1,1,1,13,1,13,13,1,13,13,1,13,13,1,13},
		// },
	} {
		var f FindSumPairs
		var output any
		for idx, command := range tc.commands {
			output = nil
			switch command {
			case "FindSumPairs":
				f = NewFindSumPairs(tc.values[idx][0], tc.values[idx][1])
			case "add":
				f.Add(tc.values[idx][0][0], tc.values[idx][0][1])
			case "count":
				output = f.Count(tc.values[idx][0][0])
			}
			assert.Equal(t, tc.output[idx], output)
		}
	}
}
