package utils

import "math"

func GenerateEnvelope(x, y, z int) map[string]float64 {
	worldMax := 20037508.3427892
	worldMin := -1 * worldMax
	worldSize := worldMax - worldMin
	worldTileSize := math.Pow(2, float64(z))
	tileSize := worldSize / worldTileSize

	envelope := make(map[string]float64)
	envelope["x_min"] = worldMin + tileSize*float64(x)
	envelope["x_max"] = worldMin + tileSize*float64(x+1)
	envelope["y_min"] = worldMax - tileSize*float64(y+1)
	envelope["y_max"] = worldMax - tileSize*float64(y)

	return envelope
}
