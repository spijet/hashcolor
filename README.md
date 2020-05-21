# Hashcolor

Simple helper for various terminal color-coding purposes.

## Install

	go get -u github.com/spijet/hashcolor/cmd/hashcolor
	
## Examples
### Colored SSH
The `examples/cssh` script provides an example of per-host color-coded SSH sessions. To use it, copy it to somewhere in `$PATH`, mark it executable and (optionally) make an alias for it:

``` sh
copy examples/cssh $HOME/.local/bin/
chmod +x $HOME/.local/bin/cssh
echo "alias ssh='cssh'" >> $HOME/.bashrc
```

`cssh` sets color code based on:
* SSH user name;
* Remote hostname/port;
* Local machine hostname.

If you run it from a shell (as opposed to running it your favourite terminal emulator's only process), it will try to reset the terminal colours back to defaults once SSH session is finished, but make sure to set your theme-provided colours in `DEFAULT_BG` and `DEFAULT_FG` variables in script.
