package app

import (
	"fmt"
	"skillbox_final/config"
)

func Run(cfg *config.Config) {

	fmt.Println(cfg)

}

func parseCSV(filename string, num_fields int) (result string, err error) {
	fmt.Println(num_fields)
	return filename, nil
}
