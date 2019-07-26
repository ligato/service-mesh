//nolint
package pods

import (
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KernelDataplanePod creates a pod
func KernelDataplanePod(name string, node *v1.Node) *v1.Pod {
	return createKernelDataplanePod(name, node, nil, nil, nil)
}

// KernelDataplanePodConfig creates a pod with config
func KernelDataplanePodConfig(name string, node *v1.Node, variables map[string]string) *v1.Pod {
	return createKernelDataplanePod(name, node, nil, nil, variables)
}

// KernelDataplanePodLiveCheck creates a pod with live check
func KernelDataplanePodLiveCheck(name string, node *v1.Node) *v1.Pod {
	return createKernelDataplanePod(name, node, createProbe("/liveness"), createProbe("/readiness"), nil)
}

func createKernelDataplanePod(name string, node *v1.Node, liveness, readiness *v1.Probe, variables map[string]string) *v1.Pod {
	ht := new(v1.HostPathType)
	*ht = v1.HostPathDirectoryOrCreate

	priv := true

	mode := v1.MountPropagationBidirectional
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name: name,
		},
		TypeMeta: v12.TypeMeta{
			Kind: "Deployment",
		},
		Spec: v1.PodSpec{
			HostPID: true,
			Volumes: []v1.Volume{
				{
					Name: "workspace",
					VolumeSource: v1.VolumeSource{
						HostPath: &v1.HostPathVolumeSource{
							Type: ht,
							Path: "/var/lib/networkservicemesh",
						},
					},
				},
				{
					Name: "postmortem",
					VolumeSource: v1.VolumeSource{
						HostPath: &v1.HostPathVolumeSource{
							Type: ht,
							Path: "/var/tmp/nsm-postmortem",
						},
					},
				},
			},
			Containers: []v1.Container{
				containerMod(&v1.Container{
					Name:            "kernel-forwarder",
					Image:           "networkservicemesh/kernel-forwarder:latest",
					ImagePullPolicy: v1.PullIfNotPresent,
					VolumeMounts: []v1.VolumeMount{
						{
							Name:             "workspace",
							MountPath:        "/var/lib/networkservicemesh/",
							MountPropagation: &mode,
						},
						{
							Name:             "postmortem",
							MountPath:        "/var/tmp/nsm-postmortem/",
							MountPropagation: &mode,
						},
					},
					Env: []v1.EnvVar{
						{
							Name: "NSM_DATAPLANE_SRC_IP",
							ValueFrom: &v1.EnvVarSource{
								FieldRef: &v1.ObjectFieldSelector{
									FieldPath: "status.podIP",
								},
							},
						},
						{
							Name:  "INITIAL_LOGLVL",
							Value: "debug",
						},
					},
					SecurityContext: &v1.SecurityContext{
						Privileged: &priv,
					},
					LivenessProbe:  liveness,
					ReadinessProbe: readiness,
					Resources:      createDefaultDataplaneResources(),
				}),
			},
			TerminationGracePeriodSeconds: &ZeroGraceTimeout,
		},
	}
	if len(variables) > 0 {
		for k, v := range variables {
			pod.Spec.Containers[0].Env = append(pod.Spec.Containers[0].Env, v1.EnvVar{
				Name:  k,
				Value: v,
			})
		}
	}
	if node != nil {
		pod.Spec.NodeSelector = map[string]string{
			"kubernetes.io/hostname": node.Labels["kubernetes.io/hostname"],
		}
	}
	return pod
}
