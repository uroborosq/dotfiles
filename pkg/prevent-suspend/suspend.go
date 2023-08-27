package preventsuspend

import (
	"github.com/leberKleber/go-mpris"
)

func IsPlaying() (bool, error) {
	mpris.PlaybackStatus()
}
