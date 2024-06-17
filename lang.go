package lang

import (
	"errors"

	"github.com/gocarina/gocsv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
)

func InjectYamlData(tag language.Tag, data []byte) error {
	var mm []*MessageYaml
	if err := yaml.Unmarshal(data, &mm); err != nil {
		return err
	}
	for _, v := range mm {
		if err := v.Inject(tag); err != nil {
			return err
		}
	}
	return nil
}
func InjectCsvData(tag language.Tag, data []byte) error {
	var mm []*MessageCsv
	if err := gocsv.UnmarshalBytes(data, &mm); err != nil {
		return err
	}
	for _, v := range mm {
		if err := v.Inject(tag); err != nil {
			return err
		}
	}
	return nil
}

//inject data

var lang = language.Chinese
var default_printer = message.NewPrinter(lang)

func SetLang(tag language.Tag) {
	lang = tag
	default_printer = message.NewPrinter(lang)
}

func GetLang() language.Tag {
	return lang
}

func Get(s string) string {
	return default_printer.Sprint(s)
}

func GetByLang(lang language.Tag, s string) string {
	p := message.NewPrinter(lang)
	return p.Sprint(s)
}

func Getf(s string, args ...any) string {
	return default_printer.Sprintf(s, args...)
}

func GetfByLang(lang language.Tag, s string, args ...any) string {
	p := message.NewPrinter(lang)
	return p.Sprintf(s, args...)
}

func Err(s string) error {
	return errors.New(Get(s))
}

func ErrByLang(lang language.Tag, s string) error {
	return errors.New(GetByLang(lang, s))
}

func Errf(s string, args ...any) error {
	return errors.New(Getf(s, args...))
}

func ErrfByLang(lang language.Tag, s string, args ...any) error {
	return errors.New(GetfByLang(lang, s, args...))
}
