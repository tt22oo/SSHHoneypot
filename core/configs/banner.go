package configs

import "os"

func ReadBanner() ([]byte, error) {
	data, err := os.ReadFile("configs/banner.txt")
	if err != nil {
		return nil, err
	}
	return data, nil
}
