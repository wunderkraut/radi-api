package setting

import (
	"github.com/wunderkraut/radi-api/property"
	"github.com/wunderkraut/radi-api/usage"
)

/**
 * Here are the commond shared Properties for the various
 * Setting operations.
 */

const (
	// property id for a single setting key
	OPERATION_PROPERTY_SETTING_KEY = "setting.key"
	// property id for scope for a single setting
	OPERATION_PROPERTY_SETTING_SCOPE = "setting.scope"
	// property id for an orerered list of keys
	OPERATION_PROPERTY_SETTING_KEYS = "setting.keys"
	// property id for a single setting value (string)
	OPERATION_PROPERTY_SETTING_VALUE = "setting.value"
)

// Property for a single setting key
type SettingKeyProperty struct {
	property.StringProperty
}

// Id for the Property
func (key *SettingKeyProperty) Id() string {
	return OPERATION_PROPERTY_SETTING_KEY
}

// Label for the Property
func (key *SettingKeyProperty) Label() string {
	return "Setting key."
}

// Description for the Property
func (key *SettingKeyProperty) Description() string {
	return "Setting string key."
}

// Is the Property internal only
func (key *SettingKeyProperty) Usage() usage.Usage {
	return property.Usage_Required()
}

// Property for a single setting scope
type SettingScopeProperty struct {
	property.StringProperty
}

// Id for the Property
func (scope *SettingScopeProperty) Id() string {
	return OPERATION_PROPERTY_SETTING_SCOPE
}

// Label for the Property
func (scope *SettingScopeProperty) Label() string {
	return "Setting scope."
}

// Description for the Property
func (scope *SettingScopeProperty) Description() string {
	return "Property string scope."
}

// Is the Property internal only
func (scope *SettingScopeProperty) Usage() usage.Usage {
	return property.Usage_Optional()
}

// Property for an ordered list of config keys
type SettingKeysProperty struct {
	property.StringSliceProperty
}

// Id for the Property
func (keys *SettingKeysProperty) Id() string {
	return OPERATION_PROPERTY_SETTING_KEYS
}

// Label for the Property
func (keys *SettingKeysProperty) Label() string {
	return "Property key list."
}

// Description for the Property
func (keys *SettingKeysProperty) Description() string {
	return "Property key list."
}

// Is the Property internal only
func (keys *SettingKeysProperty) Usage() usage.Usage {
	return property.Usage_ReadOnly()
}

// Property for a single config value
type SettingValueProperty struct {
	property.BytesArrayProperty
}

// Id for the Property
func (settingValue *SettingValueProperty) Id() string {
	return OPERATION_PROPERTY_SETTING_VALUE
}

// Label for the Property
func (settingValue *SettingValueProperty) Label() string {
	return "Property value."
}

// Description for the Property
func (settingValue *SettingValueProperty) Description() string {
	return "Property value."
}

// Is the Property internal only
func (settingValue *SettingValueProperty) Usage() usage.Usage {
	return property.Usage_ReadOnly()
}
