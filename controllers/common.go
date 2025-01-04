/*

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

package controllers

import (
	dspav1 "github.com/opendatahub-io/data-science-pipelines-operator/api/v1"
)

var commonTemplatesDir = "common/default"

const commonCusterRolebindingTemplate = "common/no-owner/clusterrolebinding.yaml.tmpl"

func (r *DSPAReconciler) ReconcileCommon(dsp *dspav1.DataSciencePipelinesApplication, params *DSPAParams) (status string, err error) {
	log := r.Log.WithValues("namespace", dsp.Namespace).WithValues("dspa_name", dsp.Name)

	log.Info("Applying Common Resources")
	err = r.ApplyDir(dsp, params, commonTemplatesDir)
	if err != nil {
		return "Error Applying Common Resources", err
	}
	err = r.ApplyWithoutOwner(params, commonCusterRolebindingTemplate)
	if err != nil {
		return "Error Applying clusterrolebinding", err
	}

	log.Info("Finished applying Common Resources")
	return "Common Resources Applied", nil
}

func (r *DSPAReconciler) CleanUpCommon(params *DSPAParams) error {
	err := r.DeleteResource(params, commonCusterRolebindingTemplate)
	if err != nil {
		return err
	}
	return nil
}
