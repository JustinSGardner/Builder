package workspace

import (
	"fmt"
	"log"
	"os"
)

func workSpaceDir(path string) (bool, error) {
	//check if file path exists, returns err = nil if file exists
	_, err := os.Stat(path)

	if err == nil {
		fmt.Println("Path already exists")
	}

	// should return true if file doesn't exist
	if os.IsNotExist(err) {

		errDir := os.MkdirAll(path, 0755)
		//should return nil once directory is made, if not, throw err
		if errDir != nil {
			log.Fatal(err)
		}
		fmt.Println("Path created")
	}

	//check workspace env exists, if not, create it
	val, present := os.LookupEnv("BUILDER_WORKSPACE_DIR")
	if !present {
		os.Setenv("BUILDER_WORKSPACE_DIR", path)
		fmt.Println("BUILDER_WORKSPACE_DIR", os.Getenv("BUILDER_WORKSPACE_DIR"))
	} else {
		fmt.Println("BUILDER_WORKSPACE_DIR", val)
	}
	return true, err
}

//MakeWorkspaceDir does...
func MakeWorkspaceDir(path string) {

	workPath := path + "/workspace"

	fmt.Printf(workPath)
	workSpaceDir(workPath)

}
