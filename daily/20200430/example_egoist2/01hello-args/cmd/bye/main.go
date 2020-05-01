package main
import (
	"flag"
	"fmt"
	"os"
	"log"
)

// this package is generated by egoist.generate.clikit

// Option ...
type Option struct {
	Name string // for `-name`
	Args []string // cmd.Args
}


func main()  {
	opt := &Option{}
	cmd := flag.NewFlagSet("bye", flag.ContinueOnError)
	cmd.Usage = func(){
		fmt.Fprintln(cmd.Output(), `bye - bye message`)
		fmt.Fprintln(cmd.Output(), "")
		fmt.Fprintln(cmd.Output(), "Usage:")
		cmd.PrintDefaults()
	}
	cmd.StringVar(&opt.Name, "name", "", "-")

	if err := cmd.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			cmd.Usage()
		}
		os.Exit(1)
	}
	opt.Args = cmd.Args()
	if err := run(opt); err != nil {
		log.Fatalf("!!%+v", err)
	}
}

func run(opt *Option) error {
	for _, target := range opt.Args  {
		fmt.Printf("%s: bye %s\n", opt.Name, target)
	}
	return nil
}