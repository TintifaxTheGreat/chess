# Chess

The subject of my current live stream series.
A simple chess UI built with [Fyne](https://fyne.io).

Thanks to:

* Chess Go library by Logan Spears github.com/notnil/chess
* Pieces created by Cburnett, accessed from Wikipedia with BSD license.
(for example https://commons.wikimedia.org/wiki/File:Chess_kdt45.svg)


## Running

Just use the go tools to install on your system and run it.

    $ go get github.com/andydotxyz/chess
    $ ./chess

## Installing

To install alongside the other applications on your system use the `fyne` tool.

    $ go get fyne.io/fyne/v2/cmd/fyne
    $ fyne install

## Screenshot

![](/img/screenshot.png)

## Status

- [x] Renders board
- [x] Animate moves
- [x] Polish board and colours etc
- [x] Handle user input
- [x] Drag and drop for moves
- [X] Take turns against a computer player
- [X] Save state and restore on app launch

TODO

- [ ] Add game summary info (who to move etc)
- [ ] Remove dependency on external algorithm
