package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SecretTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Data               interface{} `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`

	// stringData allows specifying non-binary secret data in string form.
	// It is provided as a write-only convenience method.
	// All keys and values are merged into the data field on write, overwriting any existing values.
	// It is never output when reading from the API.
	// +k8s:conversion-gen=false
	// +optional
	StringData interface{} `json:"stringData,omitempty" protobuf:"bytes,4,rep,name=stringData"`

	// Used to facilitate programmatic handling of secret data.
	// +optional
	Type string `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType"`
}
