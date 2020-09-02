package gallon

import (
	"strconv"
	"strings"

	"github.com/RafilxTenfen/go-gallon/gallons"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/log"
)

func (t *terminalHandler) run(cmd *libcmd.Cmd) error {
	w := cmd.GetUint("water")
	gallonsStr := cmd.GetString("gallon")
	strGallons := strings.Split(*gallonsStr, ",")
	volumes := make([]uint32, len(strGallons))

	for i, str := range strGallons {
		volume, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		volumes[i] = uint32(volume)
	}

	watter := uint32(*w)
	group := gallons.CreateBottleGroup(watter, volumes...)

	restWatter, bottlesUsed, err := group.UseBottles()
	if err != nil {
		log.WithError(err).Error("Error on Use Bottles")
		return err
	}

	log.With(log.F{
		"Total of Watter(ML)": watter,
		"Rest of Watter(ML)":  restWatter,
		"Bottles(ML)":         bottlesUsed,
	}).Info("The bottles to be used to fill the watter are")
	return nil
}
