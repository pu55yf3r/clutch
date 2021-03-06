package k8s

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"

	k8sapiv1 "github.com/lyft/clutch/backend/api/k8s/v1"
)

func (s *svc) DescribeDeployment(ctx context.Context, clientset, cluster, namespace, name string) (*k8sapiv1.Deployment, error) {
	cs, err := s.manager.GetK8sClientset(clientset, cluster, namespace)
	if err != nil {
		return nil, err
	}

	deployments, err := cs.AppsV1().Deployments(cs.Namespace()).List(ctx, metav1.ListOptions{
		FieldSelector: "metadata.name=" + name,
	})
	if err != nil {
		return nil, err
	}

	if len(deployments.Items) == 1 {
		return ProtoForDeployment(cs.Cluster(), &deployments.Items[0]), nil
	} else if len(deployments.Items) > 1 {
		return nil, fmt.Errorf("Located multiple Deployments")
	}

	return nil, fmt.Errorf("Unable to locate Deployment")
}

func ProtoForDeployment(cluster string, deployment *appsv1.Deployment) *k8sapiv1.Deployment {
	clusterName := deployment.ClusterName
	if clusterName == "" {
		clusterName = cluster
	}
	return &k8sapiv1.Deployment{
		Cluster:     clusterName,
		Namespace:   deployment.Namespace,
		Name:        deployment.Name,
		Labels:      deployment.Labels,
		Annotations: deployment.Annotations,
	}
}

func (s *svc) UpdateDeployment(ctx context.Context, clientset, cluster, namespace, name string, fields *k8sapiv1.UpdateDeploymentRequest_Fields) error {
	cs, err := s.manager.GetK8sClientset(clientset, cluster, namespace)
	if err != nil {
		return err
	}

	getOpts := metav1.GetOptions{}
	oldDeployment, err := cs.AppsV1().Deployments(cs.Namespace()).Get(ctx, name, getOpts)
	if err != nil {
		return err
	}

	newDeployment := oldDeployment.DeepCopy()
	mergeDeploymentLabelsAndAnnotations(newDeployment, fields)

	patchBytes, err := GenerateStrategicPatch(oldDeployment, newDeployment, appsv1.Deployment{})
	if err != nil {
		return err
	}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := cs.AppsV1().Deployments(cs.Namespace()).Patch(ctx, oldDeployment.Name, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
		return err
	})
	return retryErr
}

func (s *svc) DeleteDeployment(ctx context.Context, clientset, cluster, namespace, name string) error {
	cs, err := s.manager.GetK8sClientset(clientset, cluster, namespace)
	if err != nil {
		return err
	}

	opts := metav1.DeleteOptions{}
	return cs.AppsV1().Deployments(cs.Namespace()).Delete(ctx, name, opts)
}

func mergeDeploymentLabelsAndAnnotations(deployment *appsv1.Deployment, fields *k8sapiv1.UpdateDeploymentRequest_Fields) {
	if len(fields.Labels) > 0 {
		deployment.Labels = labels.Merge(labels.Set(deployment.Labels), labels.Set(fields.Labels))
		deployment.Spec.Template.ObjectMeta.Labels = labels.Merge(labels.Set(deployment.Spec.Template.ObjectMeta.Labels), labels.Set(fields.Labels))
	}

	if len(fields.Annotations) > 0 {
		deployment.Annotations = labels.Merge(labels.Set(deployment.Annotations), labels.Set(fields.Annotations))
		deployment.Spec.Template.ObjectMeta.Annotations = labels.Merge(labels.Set(deployment.Spec.Template.ObjectMeta.Annotations), labels.Set(fields.Annotations))
	}
}
