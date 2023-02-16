package internal

import "gorm.io/gorm"

// Правило для поля
type Rule struct {
	Key   string
	Value interface{}
}

// Набор правил
type Rules []Rule

// Прочие атрибуты поля
type Attribute string

// Набор атрибутов
type Attributes map[string]Attribute

// Поле формы
type FormField struct {
	gorm.Model
	Type       string
	Attributes *Attributes
	Rules      *Rules
	Sort       uint
}

// Поля формы
type FormFields map[int]FormField

// Абстрактная форма
type Form struct {
	// Id          uint64
	gorm.Model
	Title       string
	Description *string
	Fields      *FormFields
	Sources     *Sourses
}

// Источник. Описание страницы, где расположена форма
type Source struct {
	gorm.Model
	Id          uint64
	Name        string
	Description *string
	Link        string
}

type Sourses []Source
