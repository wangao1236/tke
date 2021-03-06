// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ChartGroup{}, func(obj interface{}) { SetObjectDefaults_ChartGroup(obj.(*ChartGroup)) })
	scheme.AddTypeDefaultingFunc(&ChartGroupList{}, func(obj interface{}) { SetObjectDefaults_ChartGroupList(obj.(*ChartGroupList)) })
	scheme.AddTypeDefaultingFunc(&ConfigMap{}, func(obj interface{}) { SetObjectDefaults_ConfigMap(obj.(*ConfigMap)) })
	scheme.AddTypeDefaultingFunc(&ConfigMapList{}, func(obj interface{}) { SetObjectDefaults_ConfigMapList(obj.(*ConfigMapList)) })
	scheme.AddTypeDefaultingFunc(&ImageNamespace{}, func(obj interface{}) { SetObjectDefaults_ImageNamespace(obj.(*ImageNamespace)) })
	scheme.AddTypeDefaultingFunc(&ImageNamespaceList{}, func(obj interface{}) { SetObjectDefaults_ImageNamespaceList(obj.(*ImageNamespaceList)) })
	scheme.AddTypeDefaultingFunc(&Namespace{}, func(obj interface{}) { SetObjectDefaults_Namespace(obj.(*Namespace)) })
	scheme.AddTypeDefaultingFunc(&NamespaceList{}, func(obj interface{}) { SetObjectDefaults_NamespaceList(obj.(*NamespaceList)) })
	scheme.AddTypeDefaultingFunc(&NsEmigration{}, func(obj interface{}) { SetObjectDefaults_NsEmigration(obj.(*NsEmigration)) })
	scheme.AddTypeDefaultingFunc(&NsEmigrationList{}, func(obj interface{}) { SetObjectDefaults_NsEmigrationList(obj.(*NsEmigrationList)) })
	scheme.AddTypeDefaultingFunc(&Project{}, func(obj interface{}) { SetObjectDefaults_Project(obj.(*Project)) })
	scheme.AddTypeDefaultingFunc(&ProjectList{}, func(obj interface{}) { SetObjectDefaults_ProjectList(obj.(*ProjectList)) })
	return nil
}

func SetObjectDefaults_ChartGroup(in *ChartGroup) {
	SetDefaults_ChartGroupStatus(&in.Status)
}

func SetObjectDefaults_ChartGroupList(in *ChartGroupList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ChartGroup(a)
	}
}

func SetObjectDefaults_ConfigMap(in *ConfigMap) {
	SetDefaults_ConfigMap(in)
}

func SetObjectDefaults_ConfigMapList(in *ConfigMapList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ConfigMap(a)
	}
}

func SetObjectDefaults_ImageNamespace(in *ImageNamespace) {
	SetDefaults_ImageNamespaceStatus(&in.Status)
}

func SetObjectDefaults_ImageNamespaceList(in *ImageNamespaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ImageNamespace(a)
	}
}

func SetObjectDefaults_Namespace(in *Namespace) {
	SetDefaults_NamespaceSpec(&in.Spec)
	SetDefaults_NamespaceStatus(&in.Status)
}

func SetObjectDefaults_NamespaceList(in *NamespaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Namespace(a)
	}
}

func SetObjectDefaults_NsEmigration(in *NsEmigration) {
	SetDefaults_NsEmigrationStatus(&in.Status)
}

func SetObjectDefaults_NsEmigrationList(in *NsEmigrationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NsEmigration(a)
	}
}

func SetObjectDefaults_Project(in *Project) {
	SetDefaults_ProjectSpec(&in.Spec)
	SetDefaults_ProjectStatus(&in.Status)
}

func SetObjectDefaults_ProjectList(in *ProjectList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Project(a)
	}
}
