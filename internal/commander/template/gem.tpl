*{{ .Title }}*

⚜ *Подземелья наказания:*
{{ range .DropInfo.WeeklyBosses -}}
{{ .Title }} ({{ .Domain }}), {{ .Location }}
{{ end }}
💥 *Мировые боссы:*
{{ range .DropInfo.WorldBosses -}}
{{ .Title }}, {{ .Location }}
{{ end }}