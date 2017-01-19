/*
Copyright 2016 The Kubernetes Authors.

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

package images

import (
	"fmt"
	"runtime"

	kubeadmapi "kubeadm/app/apis/kubeadm"
	"strings"
)

const (
	KubeEtcdImage = "etcd"

	KubeAPIServerImage         = "apiserver"
	KubeControllerManagerImage = "controller-manager"
	KubeSchedulerImage         = "scheduler"
	KubeProxyImage             = "proxy"

	KubeDNSImage         = "kube-dns"
	KubeDNSmasqImage     = "dnsmasq"
	KubeExechealthzImage = "exechealthz"
	KubeDiscoveryImage   = "kube-discovery"
	Pause                = "pause"
)

var (
	gcrPrefix   = "gcr.io/google_containers"
	etcdVersion = "2.2.5"

	kubeDNSVersion       = "1.7"
	kubeDiscoveryVersion = "1.0"
	dnsmasqVersion       = "1.3"
	exechealthzVersion   = "1.1"
	pauseVersion         = "3.0"
)

func CustomImage(imagesInfos map[string]string) {
	fmt.Println("==============>>  CustomImage: ", imagesInfos)
	if imagesInfos != nil && len(imagesInfos) > 0 {
		for k, v := range imagesInfos {
			fmt.Printf("============ key:%s === value:%s == ", k, v)
			switch strings.TrimSpace(k) {
			case "imagePrefix":
				gcrPrefix = v
			case "kubeDNSVersion":
				kubeDNSVersion = v
			case "dnsmasqVersion":
				dnsmasqVersion = dnsmasqVersion
			case "exechealthzVersion":
				exechealthzVersion = v
			case "pauseVersion":
				pauseVersion = v
			default:
				fmt.Printf("=========>>> unknow: %s    == %s ", k, v)
			}
		}
	}

	fmt.Println("==============>>  CustomImage:  gcrPrefix ", gcrPrefix)
}

func GetCoreImage(image string, cfg *kubeadmapi.MasterConfiguration, overrideImage string) string {
	if overrideImage != "" {
		return overrideImage
	}

	fmt.Println("===========>>> GetCoreImage gcrPrefix ", gcrPrefix, "cfg.KubernetesVersion", cfg.KubernetesVersion)
	return map[string]string{
		KubeEtcdImage:              fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "etcd", runtime.GOARCH, etcdVersion),
		KubeAPIServerImage:         fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-apiserver", runtime.GOARCH, cfg.KubernetesVersion),
		KubeControllerManagerImage: fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-controller-manager", runtime.GOARCH, cfg.KubernetesVersion),
		KubeSchedulerImage:         fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-scheduler", runtime.GOARCH, cfg.KubernetesVersion),
		KubeProxyImage:             fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-proxy", runtime.GOARCH, cfg.KubernetesVersion),
	}[image]
}

func GetAddonImage(image string) string {
	return map[string]string{
		KubeDNSImage:         fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kubedns", runtime.GOARCH, kubeDNSVersion),
		KubeDNSmasqImage:     fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-dnsmasq", runtime.GOARCH, dnsmasqVersion),
		KubeExechealthzImage: fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "exechealthz", runtime.GOARCH, exechealthzVersion),
		Pause:                fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "pause", runtime.GOARCH, pauseVersion),
		KubeDiscoveryImage:   fmt.Sprintf("%s/%s-%s:%s", gcrPrefix, "kube-discovery", runtime.GOARCH, kubeDiscoveryVersion),
	}[image]
}
