// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package node

import k8sConst "github.com/cilium/cilium/pkg/k8s/apis/cilium.io"

const (
	templateHostEndpointID = uint64(0xffff)
)

var (
	labels       map[string]string
	K8sPodName   string
	K8sNamespace string
	endpointID   = templateHostEndpointID
)

// GetLabels returns the labels of this node.
func GetLabels() map[string]string {
	return labels
}

// SetLabels sets the labels of this node.
func SetLabels(l map[string]string) {
	labels = l
}

// GetEndpointID returns the ID of the host endpoint for this node.
func GetEndpointID() uint64 {
	return endpointID
}

// SetLabels sets the ID of the host endpoint this node.
func SetEndpointID(id uint64) {
	endpointID = id
}

func GetNameAndNamespace() (string, string) {
	return labels[k8sConst.PodNameLabel], labels[k8sConst.PodNamespaceLabel]
}
