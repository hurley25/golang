// mplayer.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"pkg/mplayer/mlib"
	"pkg/mplayer/mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int


