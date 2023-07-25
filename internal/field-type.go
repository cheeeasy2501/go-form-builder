package internal

type FieldType int

// Идентификаторы поля, уникальные, не изменять
const (
	TextType FieldType = iota + 1
	ButtonType
	CheckBoxType
	ColorType
	DateType
	DateTimeType
	EmailType
	FileType
	HiddenType
	ImageType
	NumberType
	Password
	RadioType
	RangeType
	ResetType
	SearchType
	SubmitType
	TelType
	UrlType
)

type IField interface {
	Type() int
	Attributes() Attributes
	Attribute(key string) (Attribute, error)
	Value() interface{}
	Position() uint8
}

type Field struct {
	t    FieldType
	attr Attributes
	val  interface{}
	pos  uint8
}

// Получить идентификатор поля
func (f *Field) Type() FieldType {
	return f.t
}

// Получить все установленные атрибуты
func (f *Field) Attributes() Attributes {
	return f.attr
}

// Получить один атрибут по ключу(названию атрибута)
func (f *Field) Attribute(key string) (Attribute, error)

// Значение поля - value
func (f *Field) Value() interface{} {
	// в будущем сделать приведение типов через switch
	return f.val
}

func (f *Field) Position() uint8 {
	return f.pos
}

type IFields interface {
	AddField(f Field)
	Items() []Field
}

// Поля формы - будем сортировать
type Fields struct {
	items []Field
}

func (fs *Fields) AddField(f Field) {
	fs.items = append(fs.items, f)
}

func (fs *Fields) Items() []Field {
	return fs.items
}

func createField(t FieldType, val interface{}, attr Attributes, pos uint8) Field {
	return Field{
		t:    t,
		val:  val,
		attr: attr,
		pos:  pos,
	}
}

// Создание текстового поля
func NewTextField(val interface{}, attr Attributes, pos uint8) Field {
	return createField(TextType, val, attr, pos)
}

/*
*
Создание кнопки
val = текст кнопки
*/
func NewButtonField(val interface{}, attr Attributes, pos uint8) Field {
	return createField(ButtonType, val, attr, pos)
}
