// Created at 2021/1/24 12:01 上午
// Write by 刘万林
// Copyright AIUR TEC.

package IDNumber

import (
	"testing"
)

func TestIDNumber_IsValid(t *testing.T) {

	var IDs_True = []string{
		"11010519491231002X", //官方示例-女
		"440524188001010014", //官方示例-男
	}
	for index, id := range IDs_True {
		var id = New(id)
		if !id.IsValid() {
			t.Fatal(id)
			return
		}
		var birthday = id.GetBirthday()
		if birthday == nil {
			t.Fatal(id)
			return
		}
		var gender = id.GetGender()
		if gender == -1 {
			t.Fatal(id)
			return
		}
		switch index {
		case 0:
			if gender != Female {
				t.Fatal(id)
				return
			}
		case 1:
			if gender != Male {
				t.Fatal(id)
				return
			}
		}
	}

	var IDs_False = []string{
		"11010519491231001X", //女 - 错误数据
		"440524188001010024", //男 - 错误数据
	}
	for _, id := range IDs_False {
		var id = New(id)

		if id.IsValid() {
			t.Failed()
			return
		}
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

func TestRandomCard(t *testing.T) {
	for i := 0; i < 1000; i++ {
		cardStr := Random()
		card := New(cardStr)
		if !card.IsValid() {
			t.Fatal(cardStr)
			break
		}
	}
}
