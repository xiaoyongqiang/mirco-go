package utils

import (
	"sort"
	"strconv"
	"strings"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func SortIntAndString(a string) []string {
	var ret []string
	var afterret []string
	var beforeret []int
	msgstr := strings.TrimRight(a, ",")
	msgarr := strings.Split(msgstr, ",")
	sort.Strings(msgarr)
	for i := 0; i < len(msgarr); i++ {
		if (i > 0 && msgarr[i-1] == msgarr[i]) || len(msgarr[i]) == 0 {
			continue
		}
		k, _ := strconv.Atoi(msgarr[i])
		if k == 0 {
			afterret = append(afterret, msgarr[i])
		} else {
			beforeret = append(beforeret, k)
		}
	}
	sort.Sort(IntSlice(beforeret))
	for i := 0; i < len(beforeret); i++ {
		k := strconv.Itoa(beforeret[i])
		ret = append(ret, k)
		if i+1 == len(beforeret) {
			for _, v := range afterret {
				ret = append(ret, v)
			}
		}
	}

	return ret
}
