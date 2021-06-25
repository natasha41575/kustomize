module sigs.k8s.io/kustomize/plugin/builtin/valueaddtransformer

go 1.16

require (
	sigs.k8s.io/kustomize/api v0.8.9
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/kustomize/api => ../../../api

replace sigs.k8s.io/kustomize/kyaml => ../../../kyaml

replace gopkg.in/yaml.v3 => github.com/natasha41575/yaml v0.0.0-20210625180258-d648d2e6eae5
