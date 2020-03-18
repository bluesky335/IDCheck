package main

import (
	"fmt"
	"github.com/bluesky335/IDCheck/IdNumber"
	"github.com/bluesky335/IDCheck/USCI"
)

func main() {
	fmt.Println("社会统一信用代码校验测试")
	var USCI_TestData = map[string]string{
		"91350100M000100Y43": "官方示例",
		"91350211M0000XUF46": "厦门云上晴空航空科技有限公司", // 不符合国家标准的情况。但它是合法的代码。
		"91350100M0001TGQXM": "福建恒跃柳工机械销售有限公司", // 不符合国家标准的情况。但它是合法的代码。
		"91350203M0001FUE2P": "厦门黑脉网络科技有限公司",   // 不符合国家标准的情况。但它是合法的代码。
		"9135010557098120XJ": "福建百度博瑞网络科技有限公司",
		"91120222MA06DG3067": "天津云上畅游科技合伙企业（有限合伙）",
		"91120222MA0699U071": "天津锤子科技有限公司",
		"91310000MA1K35Y38P": "锤子科贸（上海）有限公司",
		"91110108795101314G": "谷歌信息技术（中国）有限公司",
		"914403001922038216": "华为技术有限公司",
		"91330000704202479R": "浙江小王子食品有限公司",
		"91500105MA5UM6TF71": "重庆云嘟嘟科技有限公司",
		"913504001555807674": "福建一建集团有限公司",
		"9135020026015919XW": "厦门象屿集团有限公司",
		"913502006120495420": "厦门中骏集团有限公司",
	}

	for USCI_Str, name := range USCI_TestData {

		var usci = USCI.New(USCI_Str)
		if usci.IsValid() {
			fmt.Printf("✅  正确-> %-20s : %s\n", USCI_Str, name)
		} else {
			fmt.Printf("❌  错误-> %-20s : %s\n", USCI_Str, name)
		}
	}

	fmt.Println("身份证号码校验测试")

	var IDs = []string{
		"11010519491231002X", //官方示例-女
		"440524188001010014", //官方示例-男
	}
	for index, id := range IDs {
		var id = IdNumber.New(id)
		var birthday, _ = id.GetBirthday()
		var gender, _ = id.GetGender()
		genderMap := map[IdNumber.Gender]string{
			IdNumber.Female: "女",
			IdNumber.Male:   "男",
		}
		if id.IsValid() {
			fmt.Printf("%-4d%s -> %s\t生日：%s 性别：%s \n", index, id, "✅", birthday, genderMap[gender])
		} else {
			fmt.Printf("%-4d%s -> %s\t生日：%s 性别：%s \n", index, id, "❌", birthday, genderMap[gender])
		}
	}
}
