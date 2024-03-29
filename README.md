# 身份证号码、法人和其他组织统一社会信用代码 的合法性校验

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

 中文 | [English](README-en.md)

用于校验 `身份证号码` 和 `法人和其他组织统一社会信用代码` 的工具。

计算规则参考国家标准文件：

- **标准号：GB 11643-1999**：[公民身份证号码](http://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E)

- **标准号：GB 32100-2015**：[法人和其他组织统一社会信用代码编码规则](http://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=24691C25985C1073D3A7C85629378AC0)

> 注：由于早期部分试点地区推行 `法人和其他组织统一社会信用代码` 较早，会存在部分代码不符合国家标准的情况。但它们都是合法的代码，应当另行处理。
> 例如：
>
> **福建恒跃柳工机械销售有限公司**: `91350100M0001TGQXM` 计算出的校验位是 1 和 M 不符
>
> **厦门云上晴空航空科技有限公司**:`91350211M0000XUF46` 计算出的校验位是 R 和 6 不符
>
> **厦门黑脉网络科技有限公司**`91350203M0001FUE2P` 计算出的校验位是 J 和 P 不符

## 使用方法

```shell
    go get github.com/bluesky335/IDCheck
```

- 法人和其他组织统一社会信用代码
  
```go
  import "github.com/bluesky335/IDCheck/USCI"
  
  var usci = USCI.New("91350100M000100Y43")
  if usci.IsValid() {
      fmt.Printf("✅正确\n")
  } else {
      fmt.Printf("❌错误\n")
  }
```

- 身份证号码

``` go
    import "github.com/bluesky335/IDCheck/IdNumber"

    var id = IdNumber.New("11010519491231002X")
    if id.IsValid() {
        fmt.Printf("%s -> %s\n", id, "✅正确")
    } else {
        fmt.Printf("%s -> %s\n", id, "❌错误")
    }
    
    var birthday = id.GetBirthday()
    if birthday != nil {
        fmt.Printf("生日：%s-%s-%s\n", birthday.Year, birthday.Month, birthday.Day)
    } else {
        // 不合法的身份证    
    }
    
    var gender = id.GetGender()
    if gender != -1 {
        genderMap := map[Gender]string{
            Female: "女",
            Male:   "男",
        }
        fmt.Printf("性别：%s\n", genderMap[gender])
    } else {
        // 不合法的身份证
    }

    // 产生一个随机的符合校验规则的身份证号码。注意！它虽然符合校验规则，但不一定真实存在。
    randomIDCard := IdNumber.Random()
    fmt.Printf("随机身份证：%s\n", randomIDCard)
     
```

## Swift 版本

[IDCard](https://github.com/bluesky335/IDCard)

## JavaScript 版本

ChrisDowney1996  [validators](https://github.com/ChrisDowney1996/validators)
