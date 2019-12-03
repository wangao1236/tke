/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	internalversion "tkestack.io/tke/api/client/clientset/internalversion/typed/platform/internalversion"
)

type FakePlatform struct {
	*testing.Fake
}

func (c *FakePlatform) CSIOperators() internalversion.CSIOperatorInterface {
	return &FakeCSIOperators{c}
}

func (c *FakePlatform) Clusters() internalversion.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakePlatform) ClusterAddons() internalversion.ClusterAddonInterface {
	return &FakeClusterAddons{c}
}

func (c *FakePlatform) ClusterAddonTypes() internalversion.ClusterAddonTypeInterface {
	return &FakeClusterAddonTypes{c}
}

func (c *FakePlatform) ClusterCredentials() internalversion.ClusterCredentialInterface {
	return &FakeClusterCredentials{c}
}

func (c *FakePlatform) ConfigMaps() internalversion.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

func (c *FakePlatform) CronHPAs() internalversion.CronHPAInterface {
	return &FakeCronHPAs{c}
}

func (c *FakePlatform) GPUManagers() internalversion.GPUManagerInterface {
	return &FakeGPUManagers{c}
}

func (c *FakePlatform) Helms() internalversion.HelmInterface {
	return &FakeHelms{c}
}

func (c *FakePlatform) IPAMs() internalversion.IPAMInterface {
	return &FakeIPAMs{c}
}

func (c *FakePlatform) LBCFs() internalversion.LBCFInterface {
	return &FakeLBCFs{c}
}

func (c *FakePlatform) LogCollectors() internalversion.LogCollectorInterface {
	return &FakeLogCollectors{c}
}

func (c *FakePlatform) Machines() internalversion.MachineInterface {
	return &FakeMachines{c}
}

func (c *FakePlatform) PersistentEvents() internalversion.PersistentEventInterface {
	return &FakePersistentEvents{c}
}

func (c *FakePlatform) Prometheuses() internalversion.PrometheusInterface {
	return &FakePrometheuses{c}
}

func (c *FakePlatform) Registries() internalversion.RegistryInterface {
	return &FakeRegistries{c}
}

func (c *FakePlatform) TappControllers() internalversion.TappControllerInterface {
	return &FakeTappControllers{c}
}

func (c *FakePlatform) VolumeDecorators() internalversion.VolumeDecoratorInterface {
	return &FakeVolumeDecorators{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePlatform) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}