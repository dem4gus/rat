package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/dem4gus/rat/internal/audit"
	"github.com/fatih/color"
	"github.com/google/go-github/v83/github"
	"github.com/spf13/cobra"
)

const columns = "REPOSITORY\tDEFAULT\tPROTECTED"

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
	ghaudit := audit.NewClient(ghclient, owner, name)

	report, err := ghaudit.Run(cmd.Context())
	if err != nil {
		return err
	}

	consolePrint(report)

	return nil
}

func parseArg(arg string) (string, string) {
	split := strings.Split(arg, "/")
	if len(split) != 2 {
		return "", ""
	}
	return split[0], split[1]
}

func consolePrint(report *audit.Report) {
	padding := 5
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	protected := color.RedString("no")
	if report.Protected {
		protected = color.GreenString("yes")
	}
	fmt.Fprintln(w, columns)
	fmt.Fprintf(w, "%v\t%v\t%v\n", report.FullName, report.Branch, protected)
	w.Flush()
}
