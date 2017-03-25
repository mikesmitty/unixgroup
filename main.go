package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	// Get our user and group list
	un := os.Getenv("USER")
	if un == "" {
		os.Exit(3)
	}
	gl := strings.Split(os.Getenv("GROUP"), ",")
	if len(gl) < 1 {
		os.Exit(4)
	}

	// Get user struct, primary gid, and gid (string) array
	u, err := user.Lookup(un)
	if err != nil {
		os.Exit(2)
	}
	gid, err := strconv.ParseInt(u.Gid, 10, 64)
	if err != nil {
		os.Exit(2)
	}
	gs, err := u.GroupIds()
	if err != nil {
		os.Exit(2)
	}

	// Convert gid strings to gid ints
	var gids []int64
	for _, v := range gs {
		g, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			gids = append(gids, g)
		}
	}

	// Check each of the supplied group names/gids
	for _, v := range gl {
		ggid, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			// Group is a gid
			if ggid == gid {
				// gid matches primary group
				os.Exit(0)
			} else {
				for _, v := range gids {
					if ggid == v {
						// gid matches secondary group
						os.Exit(0)
					}
				}
			}
		} else {
			// Group is a group name
			fmt.Printf("v: %s\n", v)
			g, err := user.LookupGroup(v)
			if err != nil {
				// Group is invalid
				continue
			}
			ggid, err := strconv.ParseInt(g.Gid, 10, 64)
			if err != nil {
				// Group doesn't exist or gid isn't numeric
				continue
			}

			if ggid == gid {
				// gid matches primary group
				os.Exit(0)
			} else {
				for _, v := range gids {
					if ggid == v {
						// gid matches secondary group
						os.Exit(0)
					}
				}
			}
		}
	}

	os.Exit(1)
}
