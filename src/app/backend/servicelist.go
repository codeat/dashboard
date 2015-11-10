// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/fields"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// Configuration for an app deployment.
type AppList struct {
	// Name of the application.
	Apps []App `json:"apps"`
}


// Configuration for an app deployment.
type App struct {
	// Name of the application.
	Name string `json:"name"`

	// Docker image path for the application.
	ContainerImage string `json:"containerImage"`
}

// Deploys an app based on the given configuration. The app is deployed using the given client.
func GetApps(client *client.Client) (*AppList, error) {

	list, err := client.Pods(api.NamespaceDefault).List(labels.Everything(), fields.Everything())

	if err != nil {
		return nil, err
	}

	apps := &AppList{}

	for _, element := range list.Items {
		apps.Apps = append(apps.Apps, App{
			Name: element.ObjectMeta.Name,
			ContainerImage: element.Spec.Containers[0].Name,
		})
	}




	return apps, nil
}
