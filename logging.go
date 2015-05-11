package roomgame

import (
	"mule/mylog"
)

var ErrLogFile = "data/logs/games/island.txt"
var Log = mylog.MakeErr("ISLAND: ", ErrLogFile)
var ErrF = mylog.ErrF
