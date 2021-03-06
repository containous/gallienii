package types

import "fmt"

// Configuration Global configuration.
type Configuration struct {
	Forks []ForkConfiguration `toml:"repository"`
}

// ForkConfiguration Fork configuration.
type ForkConfiguration struct {
	Base        Repo
	Fork        Repo
	Marker      Marker
	NoCheckFork bool
	Disable     bool
}

// Repo Repository model.
type Repo struct {
	Owner  string
	Name   string
	Branch string
}

func (r Repo) String() string {
	return fmt.Sprintf("%s/%s:%s", r.Owner, r.Name, r.Branch)
}

// Marker Define labels added by the bot.
type Marker struct {
	NeedResolveConflicts string
	ByBot                string
}
