package clibase

import "github.com/spf13/cobra"

func ChainedPersistentPreRunE(preRunE func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
			if parent.PersistentPreRunE != nil {
				err := parent.PersistentPreRunE(parent, args)
				if err != nil {
					return err
				}
				break
			}
		}
		return preRunE(cmd, args)
	}
}

func ChainedPersistentPreRun(preRun func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
			if parent.PersistentPreRun != nil {
				parent.PersistentPreRun(parent, args)
				break
			}
		}
		preRun(cmd, args)
	}
}

func ChainedPersistentPostRunE(postRunE func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
			if parent.PersistentPostRunE != nil {
				err := parent.PersistentPostRunE(parent, args)
				if err != nil {
					return err
				}
				break
			}
		}
		return postRunE(cmd, args)
	}
}

func ChainedPersistentPostRun(postRun func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		for parent := cmd.Parent(); parent != nil; parent = parent.Parent() {
			if parent.PersistentPostRun != nil {
				parent.PersistentPostRun(parent, args)
				break
			}
		}
		postRun(cmd, args)
	}
}
