package image

type harborV1_10_0_ImageLocator struct {
}

func (h harborV1_10_0_ImageLocator) CoreImage() *string {
	return StringToStringPtr("goharbor/harbor-core:v1.10.0")
}

func (h harborV1_10_0_ImageLocator) ChartMuseumImage() *string {
	return StringToStringPtr("goharbor/chartmuseum-photon:v0.9.0-v1.10.0")
}

func (h harborV1_10_0_ImageLocator) ClairImage() *string {
	return StringToStringPtr("goharbor/clair-photon:v2.1.1-v1.10.0")
}

func (h harborV1_10_0_ImageLocator) ClairAdapterImage() *string {
	// As it mentioned, https://github.com/goharbor/harbor-operator/blob/44ab8a074b3ebda2c94d29268a7fc823c9fe97a9/api/v1alpha1/harbor_image.go#L29
	// Use "goharbor/clair-adapter-photon:v1.0.1-v1.10.0" when possible
	return StringToStringPtr("holyhope/clair-adapter-with-config:v1.10.0")
}

func (h harborV1_10_0_ImageLocator) JobServiceImage() *string {
	return StringToStringPtr("goharbor/harbor-jobservice:v1.10.0")
}

func (h harborV1_10_0_ImageLocator) NotaryServerImage() *string {
	return StringToStringPtr("goharbor/notary-server-photon:v0.6.1-v1.10.0")
}

func (h harborV1_10_0_ImageLocator) NotarySingerImage() *string {
	return StringToStringPtr("goharbor/notary-signer-photon:v0.6.1-v1.10.0")
}

func (h harborV1_10_0_ImageLocator) NotaryDBMigratorImage() *string {
	return StringToStringPtr("jmonsinjon/notary-db-migrator:v0.6.1")
}

func (h harborV1_10_0_ImageLocator) PortalImage() *string {
	return StringToStringPtr("goharbor/harbor-portal:v1.10.0")
}

func (h harborV1_10_0_ImageLocator) RegistryImage() *string {
	return StringToStringPtr("goharbor/registry-photon:v2.7.1-patch-2819-2553-v1.10.0")
}

func (h harborV1_10_0_ImageLocator) RegistryControllerImage() *string {
	return StringToStringPtr("goharbor/harbor-registryctl:v1.10.0")
}
