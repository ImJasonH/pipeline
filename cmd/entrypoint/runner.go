package main

import (
	"context"
	"os"
	"os/exec"

	"github.com/tektoncd/pipeline/pkg/entrypoint"
)

// TODO(jasonhall): Test that original exit code is propagated and that
// stdout/stderr are collected -- needs e2e tests.

// realRunner actually runs commands.
type realRunner struct{}

var _ entrypoint.Runner = (*realRunner)(nil)

func (*realRunner) Run(ctx context.Context, args ...string) error {
	if len(args) == 0 {
		return nil
	}
	name, args := args[0], args[1:]

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err == context.Canceled {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
