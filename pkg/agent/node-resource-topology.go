/*
Copyright 2023 Intel Corporation

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

package agent

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	policyapi "github.com/intel/nri-resmgr/pkg/policy"
	nrtapi "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
)

// UpdateNrtCR updates the node's node resource topology CR using the given data.
func (a *agent) UpdateNrtCR(policy string, zones []*policyapi.TopologyZone) error {
	a.Info("updating node resource topology CR")

	if a.nrtCli == nil {
		a.Warn("no node resource topology client, can't update CR")
		return nil
	}

	cli := a.nrtCli.NodeResourceTopologies()
	ctx := context.Background()
	cr, err := cli.Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		cr = nil
		if !errors.IsNotFound(err) {
			a.Warn("failed to look up current node resource topology CR: %v", err)
		}
	}

	// delete existing CR if we got no data from policy
	// XXX TODO Deletion should be handled differently:
	//   1. add expiration timestamp to nrtapi.NodeResourceTopology
	//   2. GC CRs that are past their expiration time (for instance by NFD)
	//   3. make sure we refresh our CR (either here or preferably/easier
	//      by triggering in resmgr an updateTopologyZones() during longer
	//      periods of inactivity)
	if len(zones) == 0 {
		if cr != nil {
			err := cli.Delete(ctx, nodeName, metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return fmt.Errorf("failed to delete node resource topology CR: %w", err)
			}
		}
		return nil
	}

	// otherwise update CR if one exists
	if cr != nil {
		cr.TopologyPolicies = []string{
			policy,
		}
		cr.Zones = zonesToNrt(zones)

		_, err = cli.Update(ctx, cr, metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("failed to update node resource topology CR: %w", err)
		}

		return nil
	}

	// or create a new one
	cr = &nrtapi.NodeResourceTopology{
		ObjectMeta: metav1.ObjectMeta{
			Name: nodeName,
		},
		TopologyPolicies: []string{
			policy,
		},
		Zones: zonesToNrt(zones),
	}

	_, err = cli.Create(ctx, cr, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create node resource topology CR: %w", err)
	}

	return nil
}

func zonesToNrt(in []*policyapi.TopologyZone) nrtapi.ZoneList {
	out := nrtapi.ZoneList{}
	for _, i := range in {
		resources := nrtapi.ResourceInfoList{}
		for _, r := range i.Resources {
			resources = append(resources, nrtapi.ResourceInfo{
				Name:        r.Name,
				Capacity:    r.Capacity,
				Allocatable: r.Allocatable,
				Available:   r.Available,
			})
		}
		out = append(out, nrtapi.Zone{
			Name:       i.Name,
			Type:       i.Type,
			Parent:     i.Parent,
			Resources:  resources,
			Attributes: attributesToNrt(i.Attributes),
		})
	}
	return out
}

func attributesToNrt(in []*policyapi.ZoneAttribute) nrtapi.AttributeList {
	var out nrtapi.AttributeList
	for _, i := range in {
		out = append(out, nrtapi.AttributeInfo{
			Name:  i.Name,
			Value: i.Value,
		})
	}

	return out
}
