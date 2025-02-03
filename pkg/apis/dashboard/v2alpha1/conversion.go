package v2alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"

	common "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1"
)

func Convert_v0alpha1_Unstructured_To_v2alpha1_DashboardSpec(in *common.Unstructured, out *DashboardSpec, s conversion.Scope) error {
	return nil
}

func Convert_v2alpha1_DashboardSpec_To_v0alpha1_Unstructured(in *DashboardSpec, out *common.Unstructured, s conversion.Scope) error {
	return nil
}
