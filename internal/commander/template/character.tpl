*{{ .Title }}* {{ .Element}} *{{ .Rarity }}*★
Регион: {{ .Region }}

*Возвышение персонажа (1-90):*
*{{ .Materials.Ascension.Gem.Title }}* — {{ .Materials.Ascension.Gem.GetTotal }} шт.
💥 *{{ .Materials.Ascension.BossDrop.Title }}* — {{ .Materials.Ascension.BossDrop.GetTotal }} шт.
Можно получить: {{ .Materials.Ascension.BossDrop.Boss }}, {{ .Materials.Ascension.BossDrop.Location }}
🌺 *{{ .Materials.Ascension.LocalSpeciality.Title }}* — {{ .Materials.Ascension.LocalSpeciality.GetLocalSpecialityTotal }} шт.
🦴 *{{ .Materials.Ascension.CommonMaterial.Title }}* — {{ .Materials.Ascension.CommonMaterial.GetCommonAscensionTotal }} шт.
🧠 *«Опыт героя»* — {{ .Materials.Ascension.HeroWit.GetTotal }} шт.
💰 *Мора* — {{ .Materials.Ascension.Mora.GetAscensionTotal }}

*Возвышение талантов (1-10):*
📚 *Книги талантов:* «{{ .Materials.TalentUpgrade.TalentBook.Title }}» — {{ .Materials.TalentUpgrade.TalentBook.GetTotal }} шт.
Когда: {{ join .Materials.TalentUpgrade.TalentBook.Weekdays ", "}}
🦴 *{{ .Materials.TalentUpgrade.CommonMaterial.Title }}* — {{ .Materials.Ascension.CommonMaterial.GetCommonAscensionTotal }} шт.
⚜ *{{ .Materials.TalentUpgrade.BossDrop.Title }}* — {{ .Materials.TalentUpgrade.BossDrop.GetTotal}} шт.
Можно получить: {{ .Materials.TalentUpgrade.BossDrop.Boss }} ({{ .Materials.TalentUpgrade.BossDrop.Domain }}), {{ .Materials.TalentUpgrade.BossDrop.Location }}
👑 *{{ .Materials.TalentUpgrade.Crown.GetTitle }}* — {{ .Materials.TalentUpgrade.Crown.GetTotal }} шт.
💰 *Мора* — {{ .Materials.Ascension.Mora.GetTalentUpgradeTotal }}