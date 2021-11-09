package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	witness "gitlab.com/testifysec/witness-cli/pkg"
)

var keyPath string
var dataType string

var signCmd = &cobra.Command{
	Use:           "sign [file]",
	Short:         "Signs a file",
	Long:          "Signs a file with the provided key source and outputs the signed file to the specified destination",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          runSign,
	Args:          cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringVarP(&keyPath, "key", "k", "", "Path to the signing key")
	signCmd.Flags().StringVarP(&dataType, "datatype", "t", "https://witness.testifysec.com/policy/v0.1", "The URI reference to the type of data being signed. Defaults to the Witness policy type")
	signCmd.Flags().StringVarP(&outFilePath, "outfile", "o", "", "File to write signed data. Defaults to stdout")
}

//todo: this logic should be broken out and moved to pkg/
//we need to abstract where keys are coming from, etc
func runSign(cmd *cobra.Command, args []string) error {
	signer, err := loadSigner()
	if err != nil {
		return err
	}

	inFilePath := args[0]
	inFile, err := os.Open(inFilePath)
	if err != nil {
		return fmt.Errorf("could not open file to sign: %v", err)
	}

	outFile, err := loadOutfile()
	if err != nil {
		return err
	}

	defer outFile.Close()
	return witness.Sign(inFile, dataType, outFile, signer)
}