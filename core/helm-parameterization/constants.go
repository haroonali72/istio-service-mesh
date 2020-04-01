package helm_parameterization

const (
	NameHelmParameter = `{{ template "{{ .Name }}.fullname" . }}`

	CommandParameters = `{{ .Values.command }}`
	ArgsParameters    = `{{ .Values.Args }}`
	PortsParameters   = `{{ .Values.Ports }}`

	LabelParameter         = `{{- include "{{ .Name }}.labels" . | nindent {{ .Indent }} }}`
	AnnotationParameter    = `{{- include "{{ .Name }}.Annotation" . | nindent {{ .Indent }} }}`
	MatchSelectorParameter = `{{- include "{{ .Name }}.matchLabels" . | nindent {{ .Indent }} }}`

	LivelinessProbParameter = `\n        {{ toYaml .Values.livenessProbe | nindent 8 }}`
	ReadinessProbParameter  = `\n        {{ toYaml .Values.readinessProbe | nindent 8 }}`

	ResourcesParameter = `\n          {{ toYaml .Values.resource | nindent 10 }}`

	ReplicasHelmParameter    = "{{ .Values.replicaCount | default 1 }}"
	ImageHelmParameter       = "{{ .Values.image }}"
	ImagePullPolicyParameter = `{{ .Values.pullPolicy | default "Always" }}`
	ImagePullSecret          = `\n       {{ .Values.imagePullSecrets }}`
)

const (
	NameFunction = `
{{- define "{{ .Name }}.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- define "{{ .Name }}.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}
`
	LabelFunction = `
{{- define "{{ .Name }}.labels" -}}
#app: {{ include "cassandra.name" . }}
chart: {{ include "{{ .Name }}.chart" . }}
release: {{ .Release.Name }}
heritage: {{ .Release.Service }}
{{ .Labels }}{{- end -}}
`
)
