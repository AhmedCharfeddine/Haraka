package main

import (
	"fmt"
	"os"

	"github.com/AhmedCharfeddine/Haraka/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "haraka",
	Short: "Haraka is a tool for transliterating latin words to arabic",
	Long: "Haraka is a lightweight tool that helps convert " +
		"Latin script into Arabic script.\n" +
		"Based on Latin input, it provides real-time transliteration suggestions " +
		"for Arabic words.\n" +
		"Haraka enables you to type in Arabic with speed",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

var mappingCmd = &cobra.Command{
	Use:     "mapping",
	Aliases: []string{"map"},
	Short:   "Transliterate by mapping",
	Long: "Carry out transliteration by mapping each letter/syllable using a dictionary.\n" +
		"Example:\n\t't' maps to 'ت'\n\t'tt' maps to 'ط'\n\t'7' maps to 'ح'",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		arabicWord, err := core.Transliterate(args[0], core.MappingStrategy)
		if err != nil {
			return err
		}
		fmt.Print(arabicWord)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Transliteration failed: '%s'\n", err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(mappingCmd)

	Execute()
}
