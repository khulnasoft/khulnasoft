package config

// Status is a point in the Appliance lifecycle that an Appliance can be in.
type Status string

func (s Status) String() string {
	return string(s)
}

const (
	ConfigmapName = "sourcegraph-appliance"

	AnnotationKeyManaged             = "appliance.khulnasoft.com/managed"
	AnnotationConditions             = "appliance.khulnasoft.com/conditions"
	AnnotationKeyCurrentVersion      = "appliance.khulnasoft.com/currentVersion"
	AnnotationKeyConfigHash          = "appliance.khulnasoft.com/configHash"
	AnnotationKeyShouldTakeOwnership = "appliance.khulnasoft.com/adopted"

	// TODO set status on configmap to communicate it across reboots
	AnnotationKeyStatus = "appliance.khulnasoft.com/status"

	StatusUnknown         Status = "unknown"
	StatusInstall         Status = "install"
	StatusInstalling      Status = "installing"
	StatusUpgrading       Status = "upgrading"
	StatusWaitingForAdmin Status = "wait-for-admin"
	StatusRefresh         Status = "refresh"
	StatusMaintenance     Status = "maintenance"
)

func IsPostInstallStatus(status Status) bool {
	switch status {
	case StatusUnknown, StatusInstall, StatusInstalling, StatusWaitingForAdmin:
		return false
	}
	return true
}
