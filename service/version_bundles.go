package service

import (
	"github.com/giantswarm/versionbundle"

	clusterapiv29 "github.com/giantswarm/aws-operator/service/controller/clusterapi/v29"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v25"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v26"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v27"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v28"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v28patch1"
	"github.com/giantswarm/aws-operator/service/controller/legacy/v29"
)

// NewVersionBundles returns the array of version bundles defined for the
// operator.
func NewVersionBundles() []versionbundle.Bundle {
	var versionBundles []versionbundle.Bundle

	versionBundles = append(versionBundles, clusterapiv29.VersionBundle())

	versionBundles = append(versionBundles, v25.VersionBundle())
	versionBundles = append(versionBundles, v26.VersionBundle())
	versionBundles = append(versionBundles, v27.VersionBundle())
	versionBundles = append(versionBundles, v28.VersionBundle())
	versionBundles = append(versionBundles, v28patch1.VersionBundle())
	versionBundles = append(versionBundles, v29.VersionBundle())

	return versionBundles
}
