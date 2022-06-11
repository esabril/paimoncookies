🔔 Что, Путешественник, готов к приключениям?
Сегодня 🗓 *{{ .Weekday }}* и сегодня в Тейвате тебя ждут:

📚 *Книги на таланты*:
{{- if .SystemData.IsSunday }}
Все возможные книги во всех открытых тобой Подземельях! Ох и сложный у тебя сегодня выбор! Но Паймон здесь, чтобы помочь!
{{- else -}}
{{- range .SystemData.Regions }}
{{ .Title }}: {{ range index $.Content.TalentBooks .Title -}} «{{ .Title }}» {{- end }}
{{- end -}}
{{- end }}

🗡 *Материалы для улучшения оружия:*
{{- if .SystemData.IsSunday }}
Сегодня мы можем получить все возможные материалы! Давай выбирать вместе, куда мы сегодня отправимся?
{{- else -}}
{{- range .SystemData.Regions }}
{{ .Title }}: {{ range index $.Content.WeaponMaterials .Title -}} «{{ .Title }}» {{- if .Alias }} ({{ .Alias -}}) {{- end -}} {{- end -}}
{{- end }}
{{- end }}

Запасись смолой и вперед! А Паймон всегда будет с тобой! 💫