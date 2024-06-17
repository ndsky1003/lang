package lang

import "golang.org/x/text/language"

type Injectable interface {
	Inject(tag language.Tag) error
}
