package cmd

import (
	"fmt"
	"os"
	"retroio/commodore/d64"

	"github.com/spf13/cobra"

	"retroio/commodore"
	"retroio/commodore/d71"
	"retroio/commodore/d81"
	"retroio/commodore/t64"
	"retroio/commodore/tap"
	"retroio/storage"
)

var commodoreGeometryCmd = &cobra.Command{
	Use:   "geometry FILE",
	Short: "Read the Commodore tape file geometry",
	Long: `Read the geometry - headers and data blocks - from a Commodore emulator TAP
or T64 tape file.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		reader, err := storage.NewReaderFromFile(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		diskSize := uint32(reader.FileSize)

		mediaType := commodoreDetermineMediaType(reader.Filename)
		if mediaType == commodore.Unknown {
			fmt.Printf("unknown media type for %s", reader.Filename)
			return
		}

		var dsk commodore.Image

		switch mediaType {
		case commodore.D64:
			if dsk, err = d64.New(reader, diskSize); err != nil {
				fmt.Println(err)
				return
			}
		case commodore.D71:
			if dsk, err = d71.New(reader, diskSize); err != nil {
				fmt.Println(err)
				return
			}
		case commodore.D81:
			if dsk, err = d81.New(reader, diskSize); err != nil {
				fmt.Println(err)
				return
			}
		case commodore.T64:
			dsk = t64.New(reader)
		case commodore.TAP:
			dsk = tap.New(reader)
		default:
			// should never reach here with the geometry command
			fmt.Print("unsupported media type for this command")
			return
		}

		if err := dsk.Read(); err != nil {
			fmt.Println("Media read error!")
			fmt.Println(err)
			os.Exit(1)
		}

		dsk.DisplayGeometry()
	},
}

func init() {
	commodoreGeometryCmd.Flags().StringVarP(&commodoreMediaTypeFlag, "media", "m", "", `Media type, default: file extension`)
	commodoreCmd.AddCommand(commodoreGeometryCmd)
}
