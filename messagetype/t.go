package messagetype

type T uint8

/*
第一种

	message.SetString(language.English, "hello", "hello %v")

	//第二种
	message.Set(language.English, "%d eggs", plural.Selectf(1, "", plural.One, "a egg", plural.Other, "%d eggs"))

	//第三种
	if err := message.Set(
		language.English,
		"%d eggs %v apples",
		catalog.Var("egg", plural.Selectf(1, "", "one", "an egg", "other", " %[1]d eggs")),
		catalog.Var("apple", plural.Selectf(2, "", "one", "an apple", "other", "apples")),
		catalog.String("%[2]d ${egg} ${apple}"),
	); err != nil {
		fmt.Println(err, 1)
	}
*/
const (
	Normal T = iota //普通字符串模式
	Plural          // 复数模式,这里只能辨别1个复数,无法组合,组合使用第三种模式,占位
	Var
)
