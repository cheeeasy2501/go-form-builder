package internal

import "gorm.io/gorm"

// Правило для поля
type Rule struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// Набор правил
type Rules []Rule

// Прочие атрибуты поля
type Attribute string

// Набор атрибутов
type Attributes map[string]Attribute

// Поле формы
type Field struct {
	gorm.Model
	Type       string      `json:"type"`
	Attributes *Attributes `json:"attributes"`
	Rules      *Rules      `json:"rules"`
	Sources    *Sourses    `json:"sources"`
	Sort       uint        `json:"sort"`
}

// Поля формы
type Fields map[int]Field

// Абстрактная форма
type Form struct {
	// Id          uint64
	gorm.Model
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Fields      *Fields  `json:"fields"`
	Sources     *Sourses `json:"sources"`
}

// Источник. Описание страницы, где расположена форма
type Source struct {
	gorm.Model
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Link        string  `json:"link"`
}

type Sourses []Source
