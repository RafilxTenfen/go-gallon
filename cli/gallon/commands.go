package gallon

import (
	"github.com/RafilxTenfen/go-gallon/cli"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/fs"
	"github.com/rhizomplatform/log"
)

// terminalHandler exists only to "pass" the
// cli.Terminal along without requiring a global variable and
type terminalHandler struct {
	term cli.Terminal
}

// Main handles all command line interfaces
func Main(args []string) int {
	h := &terminalHandler{cli.Terminal{}}
	app := libcmd.NewApp("gallon", "Basic gallon problem")
	app.Options.HelpOutput = h.term.Out()
	app.Options.StrictOperands = true

	// loggin setup
	log.Setup(fs.Path("./log"), "gallon", 30, 60)
	defer log.TearDown()
	log.SetStdoutLevel(log.LevelDebug)

	app.Command("run", "Run the gallon binary searching for the best approach leaving the least possible of watter", func(cmd *libcmd.Cmd) {
		cmd.Uint("water", 'w', 3000, "Watter value in millilitre(ML)")
		cmd.String("gallon", 'g', "1500,2000,3000", "Gallons volume comma separated in millilitre(ML)")
		cmd.Run(h.run)
	})

	// root - no command specified
	app.Run(func(cmd *libcmd.Cmd) error {
		app.PrintHelp(h.term.Out())
		return nil
	})

	if err := app.ParseArgs(args); err != nil {
		h.term.Printf("%v\n", cli.UserError(err))
		return 1
	}

	return 0
}
