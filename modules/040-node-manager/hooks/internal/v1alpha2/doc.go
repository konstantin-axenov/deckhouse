/*
Copyright 2021 Flant CJSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

// +k8s:deepcopy-gen=package

//go:generate deepcopy-gen --input-dirs github.com/deckhouse/deckhouse/modules/040-node-manager/hooks/internal/v1alpha2/ -O nodegroup_generated.deepcopy --bounding-dirs github.com/deckhouse/deckhouse/modules/040-node-manager/hooks/internal/v1alpha2/ --go-header-file ../v1/boilerplate.go.txt --output-base /tmp
//go:generate cp /tmp/github.com/deckhouse/deckhouse/modules/040-node-manager/hooks/internal/v1alpha2/nodegroup_generated.deepcopy.go ./
