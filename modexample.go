package modexample

import (
	"github.com/shuang2345/modexample/ironman"
	log "github.com/sirupsen/logrus"
)

func PrintExample() {
	log.Info("Example")
	ironman.printIronMan()
}
