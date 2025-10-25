package version

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

type versionValue int

func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return strconv.FormatBool(*v == VersionEnable)
}

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolValue, err := strconv.ParseBool(s)

	if boolValue {
		*v = VersionEnable
	} else {
		*v = VersionNotSet
	}
	return err
}

func (v versionValue) Type() string {
	return "version"
}

const (
	VersionNotSet versionValue = 0
	VersionEnable versionValue = 1
	VersionRaw    versionValue = 2
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionNotSet, "Print version information and quit.")

func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	pflag.Lookup(name).NoOptDefVal = "true"
}

func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%s\n", Get().String())
		os.Exit(0)
	} else if *versionFlag == VersionEnable {
		fmt.Printf("%s\n", Get().String())
		os.Exit(0)
	}
}
