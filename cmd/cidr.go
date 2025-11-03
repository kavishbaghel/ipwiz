/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"inet.af/netaddr"
)

// cidrCmd represents the cidr command
var cidrCmd = &cobra.Command{
	Use:   "cidr",
	Short: "Generate CIDR for a given range (First IP Address and Last IP Address)",
	Long:  `Generate CIDR for a given range (First IP Address and Last IP Address). Returns CIDR range and Number of Hosts in CIDR.`,
	Run: func(cmd *cobra.Command, args []string) {
		startIp, err := netaddr.ParseIP(args[0])

		if err != nil {
			fmt.Printf("Could not parse start IP address: %v", err)
			panic(err)
		}

		endIp, err := netaddr.ParseIP(args[1])

		if err != nil {
			fmt.Printf("Could not parse end IP address: %v", err)
			panic(err)
		}

		if startIp.Is4() != endIp.Is4() {
			panic("The start and end IPs should of the same family (IPv4 or IPv6)")
		}

	},
}

func init() {
	rootCmd.AddCommand(cidrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cidrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cidrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
