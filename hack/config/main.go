package main

import (
	"io/ioutil"
	"time"

	"github.com/appscode/go/log"
	"github.com/appscode/go/runtime"
	apis "github.com/appscode/kubed/pkg/apis/kubed/v1alpha1"
	"github.com/ghodss/yaml"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	cfg := CreateClusterConfig()
	cfg.Save(runtime.GOPath() + "/src/github.com/appscode/kubed/hack/deploy/config.yaml")

	cfgBytes, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	cfgmap := core.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubed-config",
			Namespace: "kube-system",
			Labels: map[string]string{
				"app": "kubed",
			},
		},
		Data: map[string][]byte{
			"config.yaml": cfgBytes,
		},
	}
	bytes, err := yaml.Marshal(cfgmap)
	if err != nil {
		log.Fatalln(err)
	}
	p := runtime.GOPath() + "/src/github.com/appscode/kubed/hack/deploy/kubed-config.yaml"
	ioutil.WriteFile(p, bytes, 0644)
}

func CreateClusterConfig() apis.ClusterConfig {
	return apis.ClusterConfig{
		ClusterName:        "unicorn",
		EnableConfigSyncer: true,
		EventForwarder: &apis.EventForwarderSpec{
			Rules: []apis.PolicyRule{
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "",
							Resources: []string{
								"events",
							},
						},
					},
					Namespaces: []string{"kube-system"},
				},
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "",
							Resources: []string{
								"nodes",
								"persistentvolumes",
								"persistentvolumeclaims",
							},
						},
					},
				},
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "storage.k8s.io",
							Resources: []string{
								"storageclasses",
							},
						},
					},
				},
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "extensions",
							Resources: []string{
								"ingresses",
							},
						},
					},
				},
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "voyager.appscode.com",
							Resources: []string{
								"ingresses",
							},
						},
					},
				},
				{
					Operations: []apis.Operation{apis.Create},
					Resources: []apis.GroupResources{
						{
							Group: "certificates.k8s.io",
							Resources: []string{
								"certificatesigningrequests",
							},
						},
					},
				},
			},
			Receivers: []apis.Receiver{
				{
					To:       []string{"ops@example.com"},
					Notifier: "Mailgun",
				},
			},
		},
		NotifierSecretName: "notifier-config",
		RecycleBin: &apis.RecycleBinSpec{
			Path:          "/tmp/kubed/trash",
			TTL:           metav1.Duration{Duration: 7 * 24 * time.Hour},
			HandleUpdates: false,
		},
		//Snapshotter: &api.SnapshotSpec{
		//	Schedule: "@every 6h",
		//	Sanitize: true,
		//	Backend: api.Backend{
		//		StorageSecretName: "snap-secret",
		//		GCS: &api.GCSSpec{
		//			Bucket: "restic",
		//			Prefix: "minikube",
		//		},
		//	},
		//},
		//Janitors: []api.JanitorSpec{
		//	{
		//		Kind: api.JanitorElasticsearch,
		//		TTL:  metav1.Duration{Duration: 90 * 24 * time.Hour},
		//		Elasticsearch: &api.ElasticsearchSpec{
		//			Endpoint:       "https://elasticsearch-logging.kube-system:9200",
		//			LogIndexPrefix: "logstash-",
		//			SecretName:     "elasticsearch-logging-cert",
		//		},
		//	},
		//	{
		//		Kind: api.JanitorInfluxDB,
		//		TTL:  metav1.Duration{Duration: 90 * 24 * time.Hour},
		//		InfluxDB: &api.InfluxDBSpec{
		//			Endpoint: "https://monitoring-influxdb.kube-system:8086",
		//		},
		//	},
		//},
	}
}
