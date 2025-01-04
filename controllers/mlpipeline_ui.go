/*
Copyright 2023.

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

var mlPipelineUITemplatesDir = "mlpipelines-ui"

func (r *DSPAReconciler) ReconcileUI(dsp *dspav1.DataSciencePipelinesApplication,
	params *DSPAParams) (status string, err error) {

	log := r.Log.WithValues("namespace", dsp.Namespace).WithValues("dspa_name", dsp.Name)

	if dsp.Spec.MlPipelineUI == nil || !dsp.Spec.MlPipelineUI.Deploy {
		log.Info("Skipping Application of MlPipelineUI Resources")
		return "Skipped application of MlPipelineUI Resources", nil
	}

	log.Info("Applying MlPipelineUI Resources")
	err = r.ApplyDir(dsp, params, mlPipelineUITemplatesDir)
	if err != nil {
		return "Failed to apply MlPipelineUI Resources", err
	}

	log.Info("Finished applying MlPipelineUI Resources")
	return "MlPipelineUI Resources Applied", nil
}
