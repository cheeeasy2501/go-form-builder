package internal

import "gorm.io/gorm"

// // Правило для поля
// type Rule struct {
// 	Key   string
// 	Value interface{}
// }

// // Набор правил
// type Rules []Rule


// Поле формы
// type Field struct {
// 	gorm.Model
// 	Type       string
// 	Attributes *Attributes
// 	// Rules      *Rules
// 	Sort       uint
// }

// Абстрактная форма
type Form struct {
	// Id          uint64
	gorm.Model
	Title       string
	Description *string
	Fields      *Fields
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
