package lang

import (
	"testing"

	"golang.org/x/text/language"
)

var mm_csv []*MessageCsv

func TestMessageCsv(t *testing.T) {
	str := `
	key,translation
hello1,hello288888 %v
`
	InjectCsvData(language.English, []byte(str))
	SetLang(language.English)

	if Getf("hello1", 18) != "hello288888 18" {
		t.Error("err")

	}
}
