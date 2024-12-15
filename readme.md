# Conway's Game of Life

ESP32 with OLED screen ([SSD1306](https://pkg.go.dev/tinygo.org/x/drivers/ssd1306))

[TinyGo](https://tinygo.org/) used as compiler


## VS Code
1. Install [extension](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo)

2. Click on the TinyGo status bar element at the bottom of the screen and select a target. Alternatively, you could open the command palette and search for TinyGo target.


## Execute
```sh
tinygo flash -target=nodemcu -port=/dev/tty.usbserial-0001 main.go
```