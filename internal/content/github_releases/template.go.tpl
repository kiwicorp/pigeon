{{ .Owner }}/{{ .Name }}
==========

{{- range .Releases -}}
[{{- .Name -}}] - {{- .PublishedAt -}}
=====

{{- .Description -}}

[{{- .Name -}}]: {{- .Url -}}

{{- end -}}
