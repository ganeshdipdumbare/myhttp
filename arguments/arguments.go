package arguments

import (
	"fmt"
	"log"
	"strings"
)

const (
	minArgsWithFlag      = 3 // e.g "./progName -parallel 2 google.com"(min 3)
	minArgsWithFlagEqual = 2 // e.g "./progName -parallel=2 google.com"(min 2)
	minArgsWithoutFlag   = 1 // e.g."./progName google.com"(min 1)
)

// validateArgs validates arguments according to condition
// if flag is given then min no of arguments should be >= 3
// otherwise min no of arguments should be >= 1
func validateArgs(args []string, flagName string) error {
	errInvalidArg := "number of arguments are invalid"

	lenArgs := len(args)
	if lenArgs < minArgsWithoutFlag {
		log.Println(errInvalidArg)
		return fmt.Errorf(errInvalidArg)
	}

	cmdLineArg := "-" + flagName
	if args[0] == cmdLineArg && lenArgs < minArgsWithFlag {
		log.Println(errInvalidArg)
		return fmt.Errorf(errInvalidArg)
	}

	if strings.HasPrefix(args[0], cmdLineArg+"=") && lenArgs < minArgsWithFlagEqual {
		log.Println(errInvalidArg)
		return fmt.Errorf(errInvalidArg)
	}

	return nil
}

// GetSitesFromArgs gets sites from the arguments according to
// if the flag is provided or not
func GetSitesFromArgs(args []string, flagName string) ([]string, error) {
	cmdLineArg := "-" + flagName

	err := validateArgs(args, flagName)
	if err != nil {
		return nil, err
	}

	sites := []string{}
	if args[0] == cmdLineArg {
		sites = append(sites, args[2:]...)
	} else if strings.HasPrefix(args[0], (cmdLineArg + "=")) {
		sites = append(sites, args[1:]...)
	} else {
		sites = append(sites, args...)
	}

	return sites, nil
}
