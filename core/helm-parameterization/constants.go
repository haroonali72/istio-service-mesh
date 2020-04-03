package helm_parameterization

//Spaces and next lines are added because of a reason. Don't change them
const (
	NameHelmParameter = `{{ template "{{ .Name }}.fullname" . }}`

	CommandParameters = `{{ toYaml .Values.command }}`
	ArgsParameters    = `{{ toYaml .Values.args }}`
	PortsParameters   = `{{ toYaml .Values.ports | nindent 8 }}`

	LabelParameter            = `{{- include "{{ .Name }}.labels" . | nindent {{ .Indent }} }}`
	AnnotationParameter       = `{{- include "{{ .Name }}.annotation" . | nindent {{ .Indent }} }}`
	MatchSelectorParameter    = `{{- include "{{ .Name }}.matchLabels" . | nindent {{ .Indent }} }}`
	LivelinessProbIfCondition = `{{- if .Values.prob.livenessProbe }}
       `
	LivelinessProbParameter  = `{{ toYaml .Values.prob.livenessProbe | nindent 8 }}`
	ReadinessProbIfCondition = `{{- if .Values.prob.readinessProbe }}
       `
	ReadinessProbParameter = `{{ toYaml .Values.prob.readinessProbe | nindent 8 }}`

	ResourcesParameter   = `{{ toYaml .Values.resource | nindent 10 }}`
	ResourcesIfCondition = `{{- if .Values.resource }}
       `
	ReplicasHelmParameter      = "{{ .Values.replicaCount | default 1 }}"
	ImageHelmParameter         = "{{ .Values.image.image }}"
	ImagePullPolicyParameter   = `{{ .Values.image.imagePullPolicy | default "IfNotPresent"}}`
	ImagePullSecretIfCondition = `{{- if .Values.image.imagePullSecrets }}
       `
	ImagePullSecret = `{{ toYaml .Values.image.imagePullSecrets | nindent 8 }}`

	EndParameter = `
	{{- end }}`
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
	chartFunction = `
{{- define "{{ .Name  }}.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
`
	LabelFunction = `
{{- define "{{ .Name }}.labels" -}}
release: {{ .Release.Name }}
heritage: {{ .Release.Service }}
{{ .Labels }}{{- end -}}
`
)
