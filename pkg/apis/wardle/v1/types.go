package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Whitelist2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	ID   int32    `json:"id,omitempty"`
	Ips  []string `json:"ips"`
	Name string   `json:"name,omitempty"`
}

type Whitelist2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Whitelist2 `json:"items" protobuf:"bytes,2,rep,name=items"`
}