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

package IDNumber

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type IDNumber string

func New(number string) IDNumber {
	return IDNumber(number)
}

// 检查是否符合身份证国标
// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E
func (id IDNumber) IsValid() bool {
	var idStr = strings.ToUpper(string(id))
	var reg, err = regexp.Compile(`^[0-9]{17}[0-9X]$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(idStr)) {
		return false
	}
	var signChar = subString(idStr, 17, 1)
	a1Str := getSignCodeForIDNumber(idStr)
	return a1Str == signChar
}

func getSignCodeForIDNumber(idStr string) string {

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
	var sum int
	for index, c := range idStr {
		var i = 18 - index
		if index != 17 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				//计算加权因子
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return ""
			}
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str
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

type Birthday struct {
	Year  string
	Month string
	Day   string
}

// 获取身份证的生日，如果是不合法的身份证号码则返回nil
func (id IDNumber) GetBirthday() (birthday *Birthday) {
	if !id.IsValid() {
		return nil
	}
	birthday = new(Birthday)
	birthday.Year = subString(string(id), 6, 4)
	birthday.Month = subString(string(id), 10, 2)
	birthday.Day = subString(string(id), 12, 2)
	return birthday
}

type Gender int

const (
	Female Gender = 0
	Male   Gender = 1
)

// 获取身份证的性别，男性返回1，女性返回0，如果是非法的身份证，则返回-1
func (id IDNumber) GetGender() (gender Gender) {
	if !id.IsValid() {
		return -1
	}
	numStr := subString(string(id), 14, 3)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	gender = Gender(num % 2)
	return
}
