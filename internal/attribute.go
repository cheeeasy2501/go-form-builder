package internal

type AttributeType int

const (
	Basic AttributeType = iota + 1
	Validation
)

type IAttribute interface{}

// трибуты поля
type Attribute struct {
	t        AttributeType
	key, val string
}

func CreateBasicAttribute(key string, val string) Attribute {
	return Attribute{
		t:   Basic,
		key: key,
		val: val,
	}
}

func CreateValidationAttribute(key string, val string) Attribute {
	return Attribute{
		t:   Validation,
		key: key,
		val: val,
	}
}

// Набор атрибутов
type Attributes []Attribute
