/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	client "github.com/labi-le/hyprland-ipc-client"
	"github.com/mrWinston/hyprworkspacemgr/pkg/hyprland"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    Show()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")
}

func Show() {
    hl := hyprland.NewClient()
    monitors, err := hl.Monitors() 

		fmt.Println("show called")

		// Create Gtk Application, change appID to your application domain name reversed.
		const appID = "org.gtk.example"
		application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
		// Check to make sure no errors when creating Gtk Application
		if err != nil {
			log.Fatal("Could not create application.", err)
		}
		cp, err := gtk.CssProviderNew()
		if err != nil {
			log.Fatal("Error creating css provider", err)
		}
    err = cp.LoadFromPath("style.css")
		if err != nil {
			log.Fatal("Error loading stylesheet", err)
		}


		// Application signals available
		// startup -> sets up the application when it first starts
		// activate -> shows the default first window of the application (like a new document). This corresponds to the application being launched by the desktop environment.
		// open -> opens files and shows them in a new window. This corresponds to someone trying to open a document (or documents) using the application from the file browser, or similar.
		// shutdown ->  performs shutdown tasks
		// Setup activate signal with a closure function.
		application.Connect("activate", func() {
			// Create ApplicationWindow
			appWindow, err := gtk.ApplicationWindowNew(application)
			if err != nil {
				log.Fatal("Could not create application window.", err)
			}
			// Set ApplicationWindow Properties
			appWindow.SetTitle("Basic Application.")
			appWindow.SetDefaultSize(400, 400)
			appWindow.SetResizable(false)
			grid, _ := gtk.GridNew()
			numWS := 9
			for i := 0; i < numWS; i++ {
				l, err := gtk.LabelNew("")
				if err != nil {
					log.Fatal("Couldn't create gtk widget.", err)
				}
				l.SetHExpand(true)
				l.SetVExpand(true)

				grid.Attach(l, (i/3)+1, (i%3)+1, 1, 1)
        scon, _ := l.GetStyleContext()
        scon.AddProvider(cp, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
        lo.ForEach(monitors, func(item client.Monitor, index int) {
          if item.ActiveWorkspace.Id == i {
            scon.AddClass("active")
          }
        })

        grid.SetHExpand(true)
        grid.SetVExpand(true)
        grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
			}
			grid.SetRowSpacing(2)
			grid.SetColumnSpacing(2)
			grid.SetBorderWidth(2)
			appWindow.Add(grid)
			appWindow.ShowAll()
		})

		// Run Gtk application
		go func() {
			application.Run([]string{})
		}()
		time.Sleep(400* time.Millisecond)
		application.Quit()
}
