package configs

import "os"

const bannerPATH string = "configs/banner.txt"

func ReadBanner() ([]byte, error) {
	data, err := os.ReadFile(bannerPATH)
	if err != nil {
		return nil, err
	}
	return data, nil
}
