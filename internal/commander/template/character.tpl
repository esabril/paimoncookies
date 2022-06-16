*{{ .Rarity }}⭐ {{ .Element}} {{ .Title }}*
🌏 Регион: {{ .Region }}

*Материалы для прокачки:*
📚 Книги талантов: {{ .Materials.TalentBook.Title }} (🏛 {{ .Materials.TalentBook.Location }})
Можно получить в дни: {{ join .Materials.TalentBook.Weekdays ", "}}