package commands

func Execute() {
	var rootCmd = &Command{
		Name:	"paradox",
	}

	var serverCmd = &Command{
		Name: 	"server",
	}

	var configCmd = &Command{
		Name:	"config",
	}

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(configCmd)

	print((*rootCmd.commands[0]).Name)
	print((*rootCmd.commands[1]).Name)
	print("\n")
}