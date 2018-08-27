package commands

import (
    "strings"

    "github.com/urfave/cli"
)


// Validate a command's arguments
func ValidateArgs(c *cli.Context, count int, validate func(*[]string)) error {

    // Check argument count
    if len(c.Args()) != count {
        return cli.NewExitError("USAGE:" + "\n   " + c.Command.UsageText, 1);
    }

    // Validate
    messages := make([]string, 0)
    if validate != nil {
        validate(&messages)
    }

    // Return validation error or nil
    if len(messages) > 0 {
        return cli.NewExitError(strings.Join(messages, "\n"), 1)
    }
    return nil;

}

