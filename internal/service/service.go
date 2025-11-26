package service

import "strconv"

func ParseQuery(mind, maxd, minm, maxm string) (*float64, *float64, *float64, *float64) {
	var minDist, maxDist, minMass, maxMass *float64

	if mind != "" {
		if v, err := strconv.ParseFloat(mind, 64); err == nil {
			minDist = &v
		}
	}
	if maxd != "" {
		if v, err := strconv.ParseFloat(maxd, 64); err == nil {
			maxDist = &v
		}
	}
	if minm != "" {
		if v, err := strconv.ParseFloat(minm, 64); err == nil {
			minMass = &v
		}
	}
	if maxm != "" {
		if v, err := strconv.ParseFloat(maxm, 64); err == nil {
			maxMass = &v
		}
	}

	return minDist, maxDist, minMass, maxMass
}
