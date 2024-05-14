package model

import (
	"context"

	pb "github.com/erteldg/grpchealthcheckservice/pkg/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Server struct {
	Clientset *kubernetes.Clientset
	pb.UnimplementedStatusServiceServer
}

func (s *Server) GetStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	// Get the namespaces
	namespaces, err := s.Clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var ns []*pb.Namespace

	for _, namespace := range namespaces.Items {
		// Get pods current namespace
		pods, err := s.Clientset.CoreV1().Pods(namespace.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		var ps []*pb.Pod
		for _, pod := range pods.Items {
			ps = append(ps, &pb.Pod{Name: pod.Name, Status: string(pod.Status.Phase)})
		}

		// Get services current namespace
		services, err := s.Clientset.CoreV1().Services(namespace.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		var ss []*pb.Service
		for _, service := range services.Items {
			ss = append(ss, &pb.Service{Name: service.Name})
		}

		ns = append(ns, &pb.Namespace{Name: namespace.Name, Pods: ps, Services: ss})
	}

	return &pb.StatusResponse{Namespaces: ns}, nil
}
