package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/johannfh/informatik/treeman/api"
	"github.com/johannfh/informatik/treeman/frontend"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "treeman",
	Run: runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, Golang!")
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	frontendFS, err := fs.Sub(frontend.Fsys, "build")
	if err != nil {
		slog.Error("failed to create api router", "err", err)
		os.Exit(1)
	}

	r, err := api.Server{Logger: l}.CreateRouter(frontendFS)
	if err != nil {
		slog.Error("failed to create api router", "err", err)
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(":3000", r))
}
