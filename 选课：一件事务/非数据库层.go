package main

import (
	"strings"
)

func compare(choose Choose,class Class)(int){//读取choose中的timeday和class中的day比较和choose中的timelesson和class中的lesson比较
	class_day := cuttime(class.Day)
	class_lesson := cuttime(class.Lesson)
	choose_day := cuttime(choose.Timeday)
	choose_lesson := cuttime(choose.Timelesson)

	for i,u := range class_day{
		for j,v := range choose_day{
			if u == v && class_lesson[i] == choose_lesson[j]{
				return 1
			}
		}
	}

	return 0
}

func cuttime(date string)[]string{
	arr := strings.Split(date, " # ")
	return arr
}