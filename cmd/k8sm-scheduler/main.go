/*
Copyright 2015 The Kubernetes Authors.

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

package main

import (
	"fmt"
	"os"

	"github.com/kubernetes-incubator/kube-mesos-framework/pkg/hyperkube"
	"github.com/kubernetes-incubator/kube-mesos-framework/pkg/scheduler/service"
	"github.com/spf13/pflag"

)

func main() {
	s := service.NewSchedulerServer()
	s.AddStandaloneFlags(pflag.CommandLine)

	if err := s.Run(hyperkube.Nil(), pflag.CommandLine.Args()); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
