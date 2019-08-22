# KitKit
A command line tool to tag and manage your binaries with ease. KitKit lets you tag and swap binaries onto your system path.
It was made to make it easy to swap between multiple build artifacts while keeping all of your commands the same.

## Getting Started
1) Head over to the [releases](https://github.com/sosodev/KitKit/releases) section and download the binary appropriate for your system
2) rename the downloaded binary to `kitkit`
2) Optionally, set the `$KITKIT_HOME` variable to specify where you want to kitkit to store its files.
3) Add `"$KITKIT_HOME/bin"` to your path (the default location for `$KITKIT_HOME` is `"~/.kitkit"`)
4) `./kitkit add kitkit`
5) `./kitkit set kitkit latest`

KitKit is now on your path and ready to go! :tada:
Run `kitkit -h` to get more information. :smile:
