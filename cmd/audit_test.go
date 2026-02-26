package cmd

import "testing"

func TestRunAudit(t *testing.T) {
	t.Run("error on bad input", func(t *testing.T) {
		args := []string{"foo"}
		cmd := AuditCommand()
		cmd.SetArgs(args)

		if err := cmd.Execute(); err == nil {
			t.Error("error expected but none received")
		}
	})
}
