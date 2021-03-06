/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/extensions/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackupBuckets returns a BackupBucketInformer.
	BackupBuckets() BackupBucketInformer
	// BackupEntries returns a BackupEntryInformer.
	BackupEntries() BackupEntryInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ControlPlanes returns a ControlPlaneInformer.
	ControlPlanes() ControlPlaneInformer
	// Extensions returns a ExtensionInformer.
	Extensions() ExtensionInformer
	// Infrastructures returns a InfrastructureInformer.
	Infrastructures() InfrastructureInformer
	// Networks returns a NetworkInformer.
	Networks() NetworkInformer
	// OperatingSystemConfigs returns a OperatingSystemConfigInformer.
	OperatingSystemConfigs() OperatingSystemConfigInformer
	// Workers returns a WorkerInformer.
	Workers() WorkerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackupBuckets returns a BackupBucketInformer.
func (v *version) BackupBuckets() BackupBucketInformer {
	return &backupBucketInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// BackupEntries returns a BackupEntryInformer.
func (v *version) BackupEntries() BackupEntryInformer {
	return &backupEntryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ControlPlanes returns a ControlPlaneInformer.
func (v *version) ControlPlanes() ControlPlaneInformer {
	return &controlPlaneInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Extensions returns a ExtensionInformer.
func (v *version) Extensions() ExtensionInformer {
	return &extensionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Infrastructures returns a InfrastructureInformer.
func (v *version) Infrastructures() InfrastructureInformer {
	return &infrastructureInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Networks returns a NetworkInformer.
func (v *version) Networks() NetworkInformer {
	return &networkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OperatingSystemConfigs returns a OperatingSystemConfigInformer.
func (v *version) OperatingSystemConfigs() OperatingSystemConfigInformer {
	return &operatingSystemConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Workers returns a WorkerInformer.
func (v *version) Workers() WorkerInformer {
	return &workerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
