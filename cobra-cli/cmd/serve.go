/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	port int
	host string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the application server",
	Long: `The serve command will start a web server that can handle requests
   and provide API endpoints for your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if port < 1 || port > 65535 {
			fmt.Fprintf(os.Stderr, "Error: port must be between 1 and 65535\n")
			os.Exit(1)
		}

		if verbose {
			fmt.Printf("Starting server on %s:%d\n", host, port)
			fmt.Println("Verbose mode enabled")
		} else {
			fmt.Printf("Server starting on %s:%d\n", host, port)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	serveCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to bind the server to")
}
