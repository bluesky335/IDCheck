/*
MIT License

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Author: @BlueSky335 github home page : https://github.com/bluesky335

package IdNumber

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E

type IdNumber string

func New(number string) IdNumber {
	return IdNumber(number)
}

func (id IdNumber) IsValid() bool {
	//a1与对应的校验码对照表，其中key表示a1，value表示校验码，value中的10表示校验码X
	var a1Map = map[int]int{
		0:  1,
		1:  0,
		2:  10,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	var idStr = strings.ToUpper(string(id))
	var reg, err = regexp.Compile(`^[0-9]{17}[0-9X]$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(idStr)) {
		return false
	}
	var sum int
	var signChar = ""
	for index, c := range idStr {
		var i = 18 - index
		if i != 1 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				//计算加权因子
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return false
			}
		} else {
			signChar = string(c)
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str == signChar
}

type Date struct {
}

func (id IdNumber) GetBirthday() (date time.Time, err error) {
	if !id.IsValid() {
		err = errors.New("invalid id number")
		return
	}
	var yearStr = subString(string(id), 6, 4)
	var monthStr = subString(string(id), 10, 2)
	var dayStr = subString(string(id), 12, 2)
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return
	}
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return
	}
	date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return
}

type Gender int

const (
	Female Gender = 0
	Male   Gender = 1
)

func (id IdNumber) GetGender() (gender Gender, err error) {
	if !id.IsValid() {
		err = errors.New("invalid id number")
		return
	}
	numStr := subString(string(id), 14, 3)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	gender = Gender(num % 2)
	return
}

func subString(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
