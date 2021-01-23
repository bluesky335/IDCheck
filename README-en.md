# Legality check of "China Citizen ID Card Number" and "China Unified Social Credit Code for Legal Persons and Other Organizations"

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

[中文](README.md) | English

A go package for verify the ID card number and the unified social credit code of legal persons and other organizations.

Calculation rules refer to national standard documents:

- **GB 11643-1999**: [China citizen identification number](http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E)

- **GB 32100-2015**: [The coding rule of the unified social credit identifier for legal entities and other organizations](http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=24691C25985C1073D3A7C85629378AC0)

> **Note**: Due to the earlier implementation of the “Unified Social Credit Code for Legal Persons and Other Organizations” in some pilot areas of China, some codes may not meet national standards. But they are all legal codes and should be handled separately.
> For example:
>
> **福建恒跃柳工机械销售有限公司**: `91350100M0001TGQXM` The calculated check digit is `1` and `M` does not match
>
> **厦门云上晴空航空科技有限公司**: `91350211M0000XUF46` The calculated check digit is `R` and `6` does not match.
>
> **厦门黑脉网络科技有限公司**: `91350203M0001FUE2P` The calculated check digit is `J` and P dos not match.

## How to use

``` shell
go get github.com/bluesky335/IDCheck
```

- Unified social credit code for legal persons and other organizations


```go
    import "github.com/bluesky335/IDCheck/USCI"
  
    var usci = USCI.New("91350100M000100Y43")
    if usci.IsValid() {
        fmt.Printf("✅\n")
    } else {
        fmt.Printf("❌\n")
    }
```

- identification number

``` go
    import "github.com/bluesky335/IDCheck/IdNumber"

    var id = IdNumber.New("11010519491231002X")
    if id.IsValid() {
        fmt.Printf("%s -> %s\n", id, "✅")
    } else {
        fmt.Printf("%s -> %s\n", id, "❌")
    }
    
    var birthday = id.GetBirthday()
	if birthday != nil {
	    fmt.Printf("生日：%s-%s-%s\n", birthday.Year, birthday.Month, birthday.Day)
	} else {
		// invalid ID card number
	}
	
	var gender = id.GetGender()
	if gender != -1 {
	    genderMap := map[Gender]string{
			Female: "Female",
			Male:   "Male",
		}
	    fmt.Printf("性别：%s\n", genderMap[gemder])
	} else {
	    // invalid ID card number
	}
     
```

## JavaScript version

ChrisDowney1996  [validators](https://github.com/ChrisDowney1996/validators)
