package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type MessageCsv struct {
	Key         string `csv:"key"`
	Translation string `csv:"translation"`
}

func (this *MessageCsv) Inject(tag language.Tag) error {
	if this == nil {
		return nil
	}
	return message.SetString(tag, this.Key, this.Translation)
}
