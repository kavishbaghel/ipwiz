/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"inet.af/netaddr"
)

// rangeCmd represents the range command
var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Returns starting IP address, end IP address, and total host count in a given CIDR.",
	Long:  `Returns starting IP address, end IP address, and total host count in a given CIDR.".`,
	Run: func(cmd *cobra.Command, args []string) {
		CidrRange, err := netaddr.ParseIPPrefix(args[0])

		if err != nil {
			fmt.Printf("Some error occurred while fetching range: %v", err)
			panic(err)
		}

		// Calculate the start and end ip address

		networkRange := CidrRange.Range()
		startIp := networkRange.From()
		endIP := networkRange.To()

		// Calculate the host count

		bits := int(CidrRange.Bits())
		var totalBits int
		if CidrRange.IP().Is4() {
			totalBits = 32
		} else {
			totalBits = 128
		}

		hostBits := totalBits - bits
		base := 1
		for i := 0; i < hostBits; i++ {
			base = base * 2
		}
		fmt.Printf("\n Calculating range from cidr -> \n \n")
		fmt.Printf(" CIDR: %s \n \n Starting IP Address: %s \n \n Ending IP Address: %s \n \n Number of Hosts: %d \n \n", args[0], startIp, endIP, base)
	},
}

func init() {
	rootCmd.AddCommand(rangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
