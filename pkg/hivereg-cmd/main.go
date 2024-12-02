package hivereg_cmd

import (
	"flag"
	"fmt"
	go_hivereg "github.com/jc-lab/go-hivereg"
	"github.com/jc-lab/go-hivereg/model/regtype"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

type ArrayFlags []string

// String is an implementation of the flag.Value interface
func (i *ArrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Flags struct {
	Json  bool
	Store string
}

type AddFlags struct {
	*Flags
	ValueName string
	DataType  regtype.RegType
	Separator string
	Data      string
}

type commandDefine struct {
	Usage    string
	Writable bool
	Runner   func(flags *Flags, args []string, store go_hivereg.Store) error
}

var commands = map[string]commandDefine{
	"add": {
		Usage:    "add <keyname> [/v valuename] [/t datatype] [/s separator] [/d data]",
		Writable: true,
		Runner: func(flags *Flags, args []string, store go_hivereg.Store) error {
			var err error

			subFlags := &AddFlags{Flags: flags}
			if len(args) < 1 {
				return errors.New("no keyname")
			}

			keyName := args[0]
			args = args[1:]

			var dataType string
			subFlagset := flag.NewFlagSet("", flag.ExitOnError)
			subFlagset.StringVar(&subFlags.ValueName, "v", "", "")
			subFlagset.StringVar(&dataType, "t", "", "")
			subFlagset.StringVar(&subFlags.Separator, "s", "", "\x00")
			subFlagset.StringVar(&subFlags.Data, "d", "", "")
			subFlagset.Parse(args)

			if len(subFlags.ValueName) > 0 {
				subFlags.DataType, err = regtype.ValidateRegType(dataType)
				if err != nil {
					return err
				}

				err = store.AddValue(keyName, subFlags.ValueName, subFlags.DataType, subFlags.Separator, subFlags.Data)
			} else {
				err = store.AddKey(keyName)
			}

			return err
		},
	},
	"delete": {
		Usage:    "delete <keyname> [/v valuename]",
		Writable: true,
		Runner: func(flags *Flags, args []string, store go_hivereg.Store) error {
			var err error

			subFlags := &AddFlags{Flags: flags}
			if len(args) < 1 {
				return errors.New("no keyname")
			}

			keyName := args[0]
			args = args[1:]

			subFlagset := flag.NewFlagSet("", flag.ExitOnError)
			subFlagset.StringVar(&subFlags.ValueName, "v", "", "")
			subFlagset.Parse(args)

			log.Printf("keyName: %s", keyName)

			if len(subFlags.ValueName) > 0 {
				err = store.DeleteValue(keyName, subFlags.ValueName)
			} else {
				err = store.DeleteKey(keyName)
			}

			return err
		},
	},

	//"export": {
	//	Usage:    "",
	//	Writable: false,
	//	Runner: func(flags *Flags, args []string, store go_hivereg.Store) error {
	//		return nil
	//	},
	//},
	//
	//"import": {
	//	Usage: "",
	//	Writable: true,
	//	Runner: func(flags *Flags, args []string, store go_hivereg.Store) error {
	//		return nil
	//	},
	//},
}

func showUsage(program string) {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", program)

	for _, define := range commands {
		fmt.Fprintf(os.Stderr, "\t%s %s\n", program, define.Usage)
	}

	os.Exit(2)
}

func Main(args []string) {
	var flags Flags
	var err error

	program := args[0]

	flagset := flag.NewFlagSet(program, flag.ExitOnError)
	flagset.BoolVar(&flags.Json, "json", false, "Output result as JSON")
	flagset.StringVar(&flags.Store, "store", "", "Used to specify a BCD store.")

	appliedCommandFlags := make(map[string]*bool)
	for s, def := range commands {
		appliedCommandFlags[s] = flagset.Bool(s, false, def.Usage)
	}

	var fixedArgs []string
	for _, s := range args[1:] {
		if strings.Count(s, "/") == 1 && strings.HasPrefix(s, "/") {
			s = "--" + s[1:]
		}
		fixedArgs = append(fixedArgs, s)
	}
	flagset.Parse(fixedArgs)
	args = flagset.Args()

	var appliedCommand string
	for s := range commands {
		if *appliedCommandFlags[s] {
			appliedCommand = s
			break
		}
	}
	if appliedCommand == "" && len(args) >= 1 {
		appliedCommand = args[0]
		args = args[1:]
	}
	if appliedCommand == "" {
		showUsage(program)
		return
	}

	define := commands[appliedCommand]
	err = func() error {
		store, err := go_hivereg.OpenStore(flags.Store, define.Writable)
		if err != nil {
			return err
		}
		runErr := define.Runner(&flags, args, store)
		closeErr := store.Close()
		if runErr != nil {
			return runErr
		}
		return closeErr
	}()
	if err != nil {
		log.Panicln(err)
	}
}
