package storage

import (
	"context"
	"fmt"

	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	v1 "kubevirt.io/api/core/v1"

	"github.com/onsi/gomega"

	"kubevirt.io/kubevirt/tests/flags"
	"kubevirt.io/kubevirt/tests/framework/kubevirt"
	"kubevirt.io/kubevirt/tests/framework/matcher"
	"kubevirt.io/kubevirt/tests/testsuite"
)

func InitNFS(targetImage, nodeName string) *k8sv1.Pod {
	virtCli := kubevirt.Client()

	nfsPod := renderNFSServer("nfsserver", targetImage)
	nfsPod.Spec.NodeName = nodeName
	nfsPod, err := virtCli.CoreV1().Pods(testsuite.NamespacePrivileged).Create(context.Background(), nfsPod, metav1.CreateOptions{})
	gomega.Expect(err).ToNot(gomega.HaveOccurred())
	gomega.Eventually(matcher.ThisPod(nfsPod), 180).Should(matcher.BeInPhase(k8sv1.PodRunning))
	nfsPod, err = matcher.ThisPod(nfsPod)()
	gomega.Expect(err).ToNot(gomega.HaveOccurred())
	return nfsPod
}

func renderNFSServer(generateName string, hostPath string) *k8sv1.Pod {
	image := fmt.Sprintf("%s/nfs-server:%s", flags.KubeVirtRepoPrefix, flags.KubeVirtVersionTag)
	resources := k8sv1.ResourceRequirements{}
	resources.Requests = make(k8sv1.ResourceList)
	resources.Requests[k8sv1.ResourceMemory] = resource.MustParse("256M")
	resources.Requests[k8sv1.ResourceCPU] = resource.MustParse("500m")
	hostPathType := k8sv1.HostPathDirectory
	pod := &k8sv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: generateName,
			Labels: map[string]string{
				v1.AppLabel: generateName,
			},
		},
		Spec: k8sv1.PodSpec{
			RestartPolicy: k8sv1.RestartPolicyNever,
			Volumes: []k8sv1.Volume{
				{
					Name: "nfsdata",
					VolumeSource: k8sv1.VolumeSource{
						HostPath: &k8sv1.HostPathVolumeSource{
							Path: hostPath,
							Type: &hostPathType,
						},
					},
				},
			},
			Containers: []k8sv1.Container{
				{
					Name:            generateName,
					Image:           image,
					ImagePullPolicy: k8sv1.PullIfNotPresent,
					Resources:       resources,
					SecurityContext: &k8sv1.SecurityContext{
						Privileged: pointer.BoolPtr(true),
					},
					VolumeMounts: []k8sv1.VolumeMount{
						{
							Name:      "nfsdata",
							MountPath: "/data/nfs",
						},
					},
				},
			},
		},
	}
	return pod
}
