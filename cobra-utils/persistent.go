package cobra_utils

import "github.com/spf13/cobra"

func ChainPersistentE(cmd *cobra.Command, preRunE func(cmd *cobra.Command, args []string) error, postRunE func(cmd *cobra.Command, args []string) error) {
	cmd.PersistentPreRunE = ChainedPersistentPreRunE(preRunE)
	cmd.PersistentPostRunE = ChainedPersistentPostRunE(postRunE)
}

func ChainPersistent(cmd *cobra.Command, preRun func(cmd *cobra.Command, args []string), postRun func(cmd *cobra.Command, args []string)) {
	cmd.PersistentPreRun = ChainedPersistentPreRun(preRun)
	cmd.PersistentPostRun = ChainedPersistentPostRun(postRun)
}

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
		if preRunE != nil {
			return preRunE(cmd, args)
		} else {
			return nil
		}
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
		if preRun != nil {
			preRun(cmd, args)
		}
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
		if postRunE != nil {
			return postRunE(cmd, args)
		} else {
			return nil
		}
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
		if postRun != nil {
			postRun(cmd, args)
		}
	}
}
