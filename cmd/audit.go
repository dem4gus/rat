package cmd

import (
	"fmt"
	"strings"

	"github.com/dem4gus/rat/internal/audit"
	"github.com/google/go-github/v83/github"
	"github.com/spf13/cobra"
)

// AuditCommand defines the subcommand for performing audits on GitHub repositories.
// It returns a handle to the command, which can then be called in any
// other context.
func AuditCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "audit <repo>",
		Short: "Audit the settings for a GitHub repository",
		Long: `Audit scans the specified repository and reports on its
configurations.  Currently, it ensures that branch protection
is set on the default branch and that approvals are required
before merging pull requests.`,
		Args: cobra.ExactArgs(1),
		RunE: runAudit,
	}
}

func runAudit(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	owner, name := parseArg(args[0])

	if owner == "" || name == "" {
		return fmt.Errorf("could not process %q. Provide the repository as {owner}/{repo}", args[0])
	}

	ghclient := github.NewClient(nil)
	auditclient := audit.NewClient(ghclient.Repositories, owner, name)

	auditclient.Audit(cmd.Context())

	return nil
}

func parseArg(arg string) (string, string) {
	split := strings.Split(arg, "/")
	if len(split) != 2 {
		return "", ""
	}
	return split[0], split[1]
}
