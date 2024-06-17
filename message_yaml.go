package lang

import (
	"fmt"

	"github.com/ndsky1003/lang/messagetype"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"gopkg.in/yaml.v3"
)

type MessageYaml struct {
	Type    messagetype.T `yaml:"type"`
	Message yaml.Node     `yaml:"message"`
}

func (this *MessageYaml) String() string {
	return fmt.Sprintf("%+v", *this)
}

func (this *MessageYaml) Inject(tag language.Tag) error {
	if this == nil {
		return nil
	}
	var m Injectable
	switch this.Type {
	case messagetype.Normal:
		var message message_normal
		m = &message
	case messagetype.Plural:
		var message message_plural
		m = &message
	case messagetype.Var:
		var message message_var
		m = &message
	default:
		return nil
	}
	if err := this.Message.Decode(m); err != nil {
		return err
	}
	return m.Inject(tag)
}

type message_normal struct {
	Key         string `yaml:"key"`
	Translation string `yaml:"translation"`
}

func (this *message_normal) Inject(tag language.Tag) error {
	if this == nil {
		return nil
	}
	return message.SetString(tag, this.Key, this.Translation)
}

type plural_case struct {
	ArgIndex int            `yaml:"argIndex"`
	Cases    map[string]any `yaml:"cases"`
}

type message_plural struct {
	Key        string      `yaml:"key"`
	PluralCase plural_case `yaml:"pluralCase"`
}

func (this *message_plural) Inject(tag language.Tag) error {
	if this == nil {
		return nil
	}
	var cases []any
	for k, v := range this.PluralCase.Cases {
		cases = append(cases, k, v)
	}
	return message.Set(tag, this.Key, plural.Selectf(this.PluralCase.ArgIndex, "", cases...))
}

type message_placeholder struct {
	Placeholder string      `yaml:"placeholder"`
	PluralCase  plural_case `yaml:"pluralCase"`
}

type message_var struct {
	Key         string                `yaml:"key"`
	Translation string                `yaml:"translation"`
	Vars        []message_placeholder `yaml:"vars"`
}

func (this *message_var) Inject(tag language.Tag) error {
	if this == nil {
		return nil
	}
	var vars []catalog.Message
	for _, v := range this.Vars {
		var cases []any
		for k, v := range v.PluralCase.Cases {
			cases = append(cases, k, v)
		}
		vars = append(vars, catalog.Var(v.Placeholder, plural.Selectf(v.PluralCase.ArgIndex, "", cases...)))
	}
	String := catalog.String(this.Translation)
	vars = append(vars, String)

	return message.Set(tag, this.Key, vars...)
}
