package lang

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
)

func TestMain(t *testing.M) {
	str := `
- message:
      key: "hello"
      translation: "hello %v"
- type: 1
  message:
      key: "%d eggs"
      pluralCase:
          argIndex: 1
          cases:
              =1: "a egg"
              =2: "two eggs"
              other: "%d eggs"
- type: 2
  message:
      key: "%d eggs %v apples"
      translation: "${egg} %[1]d,%[2]d ${apple}"
      vars:
          - placeholder: egg
            pluralCase:
                argIndex: 1
                cases:
                    =1: "a egg"
                    =2: "two eggs"
                    other: "%[2]d eggs"
          - placeholder: apple
            pluralCase:
                argIndex: 2
                cases:
                    =1: "a apple"
                    =2: "two apples"
                    other: "%[2]d apples"

	`
	yaml.Unmarshal([]byte(str), &mm)
	for _, v := range mm {
		v.Inject(language.English)
	}
	code := t.Run()
	os.Exit(code)
}

var mm []*MessageYaml

func TestMessageYaml(t *testing.T) {
	p := message.NewPrinter(language.English)
	v := p.Sprintf("%d eggs %v apples", 1, 1)
	fmt.Println(v)
	p.Println("hello", 12)
	p.Printf("%d eggs", 3)
	fmt.Println()
}
