Fyne Render Engine
##################

..	_ResterX:	https://github.com/srwiley/rasterx
..	_Fyne:		https://github.com/fyne-io/fyne

Fyne uses the Rasterx_ rasterizer to convert path data into images that can be
*painted* on the screen. Rasterx_ accepts a Go vector of points which can
either be an open line or a closed area to be rendered.

As an example of the rasterizing process, here is a snippet of code that will render a sample path.

..  code-block::    go

    package main

    inport (
        "github.com/srwiley/rasterx"
    )

    func main() {
    }


