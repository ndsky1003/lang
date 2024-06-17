# lang
```yaml
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
```
```csv
key,translation
hello1,hello2888 %v
```
> 系统支持以上2种配置文件对应的模型如下
```go
type MessageYaml struct {
	Type    messagetype.T `yaml:"type"`
	Message yaml.Node     `yaml:"message"`
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
```

### 自定义配置文件,自定义模型
> 任意定义配置,任意模型,只需实现`Inject(tag language.Tag) error` 接口即可


#### 公开接口
```go
func InjectYamlData(tag language.Tag, data []byte) error

func InjectCsvData(tag language.Tag, data []byte) error

func SetLang(tag language.Tag)

func GetLang() language.Tag

func Get(s string) string

func GetByLang(lang language.Tag, s string) string

func Getf(s string, args ...any) string

func GetfByLang(lang language.Tag, s string, args ...any) string

func Err(s string) error

func ErrByLang(lang language.Tag, s string) error

func Errf(s string, args ...any) error

func ErrfByLang(lang language.Tag, s string, args ...any) error

```
