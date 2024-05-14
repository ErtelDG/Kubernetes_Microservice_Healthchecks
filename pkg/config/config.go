package config

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientset() (*kubernetes.Clientset, error) {
	// For dev:
	//homeDir, err := os.UserHomeDir()
	//if err != nil {
	//	log.Fatalf("Failed to get home directory: %v", err)
	//}
	//kubeconfigPath := filepath.Join(homeDir, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return clientset, nil
}
