// Created at 2021/1/24 12:01 上午
// Write by 刘万林
// Copyright AIUR TEC.

package IDNumber

import (
	"fmt"
	"testing"
)

func TestIDNumber_IsValid(t *testing.T) {

	var IDs_True = []string{
		"11010519491231002X", //官方示例-女
		"440524188001010014", //官方示例-男
	}
	for index, id := range IDs_True {
		var id = New(id)
		var birthday = id.GetBirthday()
		if birthday == nil {
			t.Failed()
			return
		}
		var gender = id.GetGender()
		if gender == -1 {
			t.Failed()
			return
		}
		genderMap := map[Gender]string{
			Female: "女",
			Male:   "男",
		}
		switch index {
		case 0:
			if gender != Female {
				t.Failed()
			}
		case 1:
			if gender != Male {
				t.Failed()
			}
		}
		if id.IsValid() {
			fmt.Printf("%-4d%s -> %s\t生日：%s-%s-%s 性别：%s \n", index, id, "✅", birthday.Year, birthday.Month, birthday.Day, genderMap[gender])
		} else {
			fmt.Printf("%-4d%s -> %s\t生日：%s-%s-%s 性别：%s \n", index, id, "❌", birthday.Year, birthday.Month, birthday.Day, genderMap[gender])
		}
	}

	var IDs_False = []string{
		"11010519491231001X", //女 - 错误数据
		"440524188001010024", //男 - 错误数据
	}
	for _, id := range IDs_False {
		var id = New(id)
		var birthday = id.GetBirthday()
		if birthday != nil {
			t.Failed()
		}
		var gender = id.GetGender()
		if gender != -1 {
			t.Failed()
		}
	}
}
