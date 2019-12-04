// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OutputSpec defines the desired state of Output
type OutputSpec struct {
	LoggingRef                string                            `json:"loggingRef,omitempty"`
	S3OutputConfig            *output.S3OutputConfig            `json:"s3,omitempty"`
	AzureStorage              *output.AzureStorage              `json:"azurestorage,omitempty"`
	GCSOutput                 *output.GCSOutput                 `json:"gcs,omitempty"`
	OSSOutput                 *output.OSSOutput                 `json:"oss,omitempty"`
	ElasticsearchOutput       *output.ElasticsearchOutput       `json:"elasticsearch,omitempty"`
	LokiOutput                *output.LokiOutput                `json:"loki,omitempty"`
	SumologicOutput           *output.SumologicOutput           `json:"sumologic,omitempty"`
	ForwardOutput             *output.ForwardOutput             `json:"forward,omitempty"`
	FileOutput                *output.FileOutputConfig          `json:"file,omitempty"`
	NullOutputConfig          *output.NullOutputConfig          `json:"nullout,omitempty"`
	KafkaOutputConfig         *output.KafkaOutputConfig         `json:"kafka,omitempty"`
	CloudWatchOutput          *output.CloudWatchOutput          `json:"cloudwatch,omitempty"`
	KinesisStreamOutputConfig *output.KinesisStreamOutputConfig `json:"kinesisStream,omitempty"`
	LumberjackOutput          *output.LumberjackOutput          `json:"lumberjack,omitempty"`
}

// OutputStatus defines the observed state of Output
type OutputStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Output is the Schema for the outputs API
type Output struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OutputSpec   `json:"spec,omitempty"`
	Status OutputStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputList contains a list of Output
type OutputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Output `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Output{}, &OutputList{})
}
