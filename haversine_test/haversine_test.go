package haversine_test

import (
	"testing"

	"github.com/umahmood/haversine"
)

var tests = []struct {
	p     Coord
	q     Coord
	outMi float64
	outKm float64
	outNm float64
}{
	{
		Coord{Lat: 22.55, Lon: 43.12},  // Rio de Janeiro, Brazil
		Coord{Lat: 13.45, Lon: 100.28}, // Bangkok, Thailand
		3786.251258825624,
		6094.544408786774,
		3294.4786344857866,
	},
	{
		Coord{Lat: 20.10, Lon: 57.30}, // Port Louis, Mauritius
		Coord{Lat: 0.57, Lon: 100.21}, // Padang, Indonesia
		3196.671009759937,
		5145.525771394785,
		2781.475296597383,
	},
	{
		Coord{Lat: 51.45, Lon: 1.15},  // Oxford, United Kingdom
		Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		863.0311907424888,
		1389.1793118293067,
		750.937437701332,
	},
	{
		Coord{Lat: 22.34, Lon: 17.05}, // Windhoek, Namibia
		Coord{Lat: 51.56, Lon: 4.29},  // Rotterdam, Netherlands
		2130.8298370015464,
		3429.89310043882,
		1854.0696039025681,
	},
	{
		Coord{Lat: 63.24, Lon: 56.59}, // Esperanza, Argentina
		Coord{Lat: 8.50, Lon: 13.14},  // Luanda, Angola
		4346.398369403186,
		6996.18595539861,
		3781.871721160945,
	},
	{
		Coord{Lat: 90.00, Lon: 0.00}, // North/South Poles
		Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		2866.1346681303867,
		4613.477506482742,
		2493.8702643424967,
	},
	{
		Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		6261.05275709582,
		10078.111954385415,
		5447.843560186316,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		mi, km, nm := haversine.Distance(input.p, input.q)

		if input.outMi != mi || input.outKm != km || input.outNm != nm {
			t.Errorf("fail: want %v %v -> %v %v got %v %v",
				input.p,
				input.q,
				input.outMi,
				input.outKm,
				mi,
				km,
				nm,
			)
		}
	}
}
