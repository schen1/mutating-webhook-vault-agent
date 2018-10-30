package webhook

import corev1 "k8s.io/api/core/v1"

type WebHook struct {
	config        *Config
	sidecarConfig *SideCarConfig
}

type Config struct {
	Template string `json:"template"`
}

type VaultAgentConfig struct {
	Config string `json:"vault-agent-config"`
}

type SideCarConfig struct {
	Containers []corev1.Container `yaml:"containers"`
	Volumes    []corev1.Volume    `yaml:"volumes"`
	ConfigMaps []corev1.ConfigMap `yaml:"configMaps"`
}

type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

type registeredAnnotation struct {
	name      string
	validator annotationValidationFunc
}

type annotationValidationFunc func(value string) error
