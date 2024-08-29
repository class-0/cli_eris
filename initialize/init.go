package initialize

import (
	"fmt"
	"os"

	// "github.com/eris-ltd/eris-cli/definitions"
	// "github.com/eris-ltd/eris-cli/perform"
	"github.com/eris-ltd/eris-cli/util"

	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

func Initialize(skipPull, verbose bool) error {

	if _, err := os.Stat(common.ErisRoot); err != nil {
		if err := common.InitErisDir(); err != nil {
			return fmt.Errorf("Could not Initialize the Eris Root Directory.\n%s\n", err)
		}
	} else {
		if verbose {
			fmt.Printf("Root eris directory (%s) already exists. Please type `eris` to see the help.\n", common.ErisRoot)
		}
	}

	if err := util.CheckDockerClient(); err != nil {
		return err
	}

	if err := InitDefaultServices(skipPull, verbose); err != nil {
		return fmt.Errorf("Could not instantiate default services.\n%s\n", err)
	}

	if verbose {
		fmt.Printf("Initialized eris root directory (%s) with default actions and service files.\n", common.ErisRoot)
	}

	// lets not pull any by default. this will spread out the get started pain
	// if skipPull {
	// 	//pull images
	// 	argsAll := []string{}
	// 	argsDef := []string{"eris/keys", "eris/ipfs", "eris/erisdb", "eris/data"}
	// 	argsDev := []string{
	// 		"erisindustries/evm_compilers",
	// 		"eris/compilers",
	// 		"erisindustries/node",
	// 		"erisindustries/python",
	// 		"erisindustries/gulp",
	// 		"erisindustries/embark_base",
	// 		"erisindustries/sunit_base",
	// 		"erisindustries/pyepm_base",
	// 	}

	// 	fmt.Println("Pulling base images...")

	// 	if !dev {
	// 		argsAll = argsDef
	// 	} else {
	// 		fmt.Println("...and development images")
	// 		argsAll = append(argsDef, argsDev...)
	// 	}

	// 	fmt.Println("This could take awhile, now is a good time to feed your marmot")
	// 	for _, img := range argsAll {
	// 		srv := definitions.BlankService()
	// 		srv.Image = img
	// 		ops := definitions.BlankOperation()
	// 		err := perform.DockerPull(srv, ops)
	// 		if err != nil {
	// 			fmt.Println("An error occured pulling the images: ", err)
	// 			return err
	// 		}
	// 	}
	// }
	// todo: when called from cli provide option to go on tour, like `ipfs tour`

	fmt.Printf("The marmots have everything set up for you.\n")
	return nil
}

func InitDefaultServices(skipPull, verbose bool) error {
	if err := dropChainDefaults(); err != nil {
		return err
	}

	if !skipPull {
		if err := pullRepo("eris-services", common.ServicesPath, verbose); err != nil {
			if verbose {
				fmt.Println("Using default defs.")
			}
			if err2 := dropDefaults(); err2 != nil {
				return fmt.Errorf("Cannot pull: %s. %s.\n", err, err2)
			}
		} else {
			if err2 := pullRepo("eris-actions", common.ActionsPath, verbose); err2 != nil {
				return fmt.Errorf("Cannot pull actions: %s.\n", err2)
			}
		}
	} else {
		if err := dropDefaults(); err != nil {
			return err
		}
	}

	return nil
}
