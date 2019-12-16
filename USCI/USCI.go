package USCI

import (
	"math"
	"regexp"
	"strings"
)

// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=24691C25985C1073D3A7C85629378AC0

//代码字符集
var charSet = []int32{
	'0',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'J',
	'K',
	'L',
	'M',
	'N',
	'P',
	'Q',
	'R',
	'T',
	'U',
	'W',
	'X',
	'Y',
}

//代码字符对应的值
var valueMap = map[int32]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'G': 16,
	'H': 17,
	'J': 18,
	'K': 19,
	'L': 20,
	'M': 21,
	'N': 22,
	'P': 23,
	'Q': 24,
	'R': 25,
	'T': 26,
	'U': 27,
	'W': 28,
	'X': 29,
	'Y': 30,
}

type item struct {
	index  int
	char   int32
	value  int
	weight int
}

type USCI string

func New(usci string) USCI {
	return USCI(usci)
}

func (usci USCI) IsValid() bool {
	var usciStr = strings.ToUpper(string(usci))
	var reg, err = regexp.Compile(`^[A-Z0-9]{18}$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(usciStr)) {
		return false
	}
	var sum = 0
	for index, c := range usciStr[:17] {
		var value = valueMap[c]
		//计算加权因子
		var weight = int(math.Pow(3, float64(index))) % 31
		sum += value * weight
	}
	var mod = sum % 31
	var sign = 31 - mod
	if sign == 31 {
		sign = 0
	}
	var signChar int32

	for key, value := range valueMap {
		signChar = key
		if value == sign {
			break
		}
	}
	var lastStr = usciStr[17:18]
	var signStr = string(signChar)
	return signStr  == lastStr
}
