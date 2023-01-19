package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var (
	port int
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server that validates the Terraform plan given through HTTP request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server up and running at port %v\n", port)
		server(port)
	},
}

func init() {
	ServerCmd.Flags().IntVarP(&port, "port", "p", 8080, "Ports the server should be listening to.")
}

func server(port int) {
	http.HandleFunc("/control", control)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func control(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[WIP] Control your plan by passing them to ptf server.\n")
}
