// +build !ignore_autogenerated

/*
Copyright 2018 The Kubed Authors.

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

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appscode/kubed/apis/kubed/v1alpha1.AzureSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"container": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"prefix": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.Backend": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"storageSecretName": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"local": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.LocalSpec"),
							},
						},
						"s3": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.S3Spec"),
							},
						},
						"gcs": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.GCSSpec"),
							},
						},
						"azure": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.AzureSpec"),
							},
						},
						"swift": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.SwiftSpec"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.AzureSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.GCSSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.LocalSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.S3Spec", "github.com/appscode/kubed/apis/kubed/v1alpha1.SwiftSpec"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.ClusterConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"clusterName": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"snapshotter": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.SnapshotSpec"),
							},
						},
						"recycleBin": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.RecycleBinSpec"),
							},
						},
						"eventForwarder": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.EventForwarderSpec"),
							},
						},
						"enableConfigSyncer": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"notifierSecretName": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"janitors": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.JanitorSpec"),
										},
									},
								},
							},
						},
						"kubeConfigFile": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"enableConfigSyncer"},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.EventForwarderSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.JanitorSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.RecycleBinSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.SnapshotSpec"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.ElasticsearchSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"endpoint": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"logIndexPrefix": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"secretName": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.EventForwarderSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"nodeAdded": {
							SchemaProps: spec.SchemaProps{
								Description: "Deprecated",
								Ref:         ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec"),
							},
						},
						"storageAdded": {
							SchemaProps: spec.SchemaProps{
								Description: "Deprecated",
								Ref:         ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec"),
							},
						},
						"ingressAdded": {
							SchemaProps: spec.SchemaProps{
								Description: "Deprecated",
								Ref:         ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec"),
							},
						},
						"warningEvents": {
							SchemaProps: spec.SchemaProps{
								Description: "Deprecated",
								Ref:         ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec"),
							},
						},
						"csrEvents": {
							SchemaProps: spec.SchemaProps{
								Description: "Deprecated",
								Ref:         ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec"),
							},
						},
						"rules": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.PolicyRule"),
										},
									},
								},
							},
						},
						"receivers": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.Receiver"),
										},
									},
								},
							},
						},
					},
					Required: []string{"rules"},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.PolicyRule", "github.com/appscode/kubed/apis/kubed/v1alpha1.Receiver"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.ForwarderSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"handle": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"namespaces": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
					Required: []string{"handle"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.GCSSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"bucket": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"prefix": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"bucket"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.GroupResources": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "GroupResources represents resource kinds in an API group.",
					Properties: map[string]spec.Schema{
						"group": {
							SchemaProps: spec.SchemaProps{
								Description: "Group is the name of the API group that contains the resources. The empty string represents the core API group.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Resources is a list of resources within the API group. Subresources are matched using a \"/\" to indicate the subresource. For example, \"pods/log\" would match request to the log subresource of pods. The top level resource does not match subresources, \"pods\" doesn't match \"pods/log\".",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
						"resourceNames": {
							SchemaProps: spec.SchemaProps{
								Description: "ResourceNames is a list of resource instance names that the policy matches. Using this field requires Resources to be specified. An empty list implies that every instance of the resource is matched.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.InfluxDBSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"endpoint": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"username": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"password": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.JanitorAuthInfo": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"CACertData": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "byte",
							},
						},
						"ClientCertData": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "byte",
							},
						},
						"ClientKeyData": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "byte",
							},
						},
						"InsecureSkipVerify": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"Username": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"Password": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"Token": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"CACertData", "ClientCertData", "ClientKeyData", "InsecureSkipVerify", "Username", "Password", "Token"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.JanitorSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"ttl": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
							},
						},
						"elasticsearch": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ElasticsearchSpec"),
							},
						},
						"influxdb": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.InfluxDBSpec"),
							},
						},
					},
					Required: []string{"kind", "ttl"},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.ElasticsearchSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.InfluxDBSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.Duration"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.KubedMetadata": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"operatorNamespace": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"searchEnabled": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
					},
					Required: []string{"searchEnabled"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.LocalSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"path": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.NoNamespacedForwarderSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"namespaces": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.PolicyRule": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"operations": {
							SchemaProps: spec.SchemaProps{
								Description: "Operation is the operation being performed",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
						"resources": {
							SchemaProps: spec.SchemaProps{
								Description: "Resources that this rule matches. An empty list implies all kinds in all API groups.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.GroupResources"),
										},
									},
								},
							},
						},
						"namespaces": {
							SchemaProps: spec.SchemaProps{
								Description: "Namespaces that this rule matches. The empty string \"\" matches non-namespaced resources. An empty list implies every namespace.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
					Required: []string{"operations"},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.GroupResources"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.Receiver": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"to": {
							SchemaProps: spec.SchemaProps{
								Description: "To whom notification will be sent",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
						"notifier": {
							SchemaProps: spec.SchemaProps{
								Description: "How this notification will be sent",
								Type:        []string{"string"},
								Format:      "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.RecycleBinSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"path": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"ttl": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
							},
						},
						"handleUpdates": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
					},
					Required: []string{"handleUpdates"},
				},
			},
			Dependencies: []string{
				"k8s.io/apimachinery/pkg/apis/meta/v1.Duration"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.ResultEntry": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"score": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"number"},
								Format: "double",
							},
						},
						"object": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
							},
						},
					},
					Required: []string{"score"},
				},
			},
			Dependencies: []string{
				"k8s.io/apimachinery/pkg/runtime.RawExtension"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.S3Spec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"endpoint": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"bucket": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"prefix": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"bucket"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.SearchResult": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
							},
						},
						"hits": {
							SchemaProps: spec.SchemaProps{
								Type: []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.ResultEntry"),
										},
									},
								},
							},
						},
						"totalHits": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int64",
							},
						},
						"maxScore": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"number"},
								Format: "double",
							},
						},
						"took": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
							},
						},
					},
					Required: []string{"totalHits", "maxScore", "took"},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.ResultEntry", "k8s.io/apimachinery/pkg/apis/meta/v1.Duration", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.SnapshotSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "For periodic full cluster backup https://github.com/appscode/kubed/issues/16",
					Properties: map[string]spec.Schema{
						"schedule": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"sanitize": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"overwrite": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"storageSecretName": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"local": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.LocalSpec"),
							},
						},
						"s3": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.S3Spec"),
							},
						},
						"gcs": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.GCSSpec"),
							},
						},
						"azure": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.AzureSpec"),
							},
						},
						"swift": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("github.com/appscode/kubed/apis/kubed/v1alpha1.SwiftSpec"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/appscode/kubed/apis/kubed/v1alpha1.AzureSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.GCSSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.LocalSpec", "github.com/appscode/kubed/apis/kubed/v1alpha1.S3Spec", "github.com/appscode/kubed/apis/kubed/v1alpha1.SwiftSpec"},
		},
		"github.com/appscode/kubed/apis/kubed/v1alpha1.SwiftSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"container": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"prefix": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
	}
}
