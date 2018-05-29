package luas

var allStops = []*Stop{
	&Stop{
		Name:        "Tallaght",
		NameAbv:     "TAL",
		Line:        "red",
		Coordinates: []float64{53.28749444, -6.374588889},
	},
	&Stop{
		Name:        "Drimnagh",
		NameAbv:     "DRI",
		Line:        "red",
		Coordinates: []float64{53.33536111, -6.318161111},
	},
	&Stop{
		Name:        "Goldenbridge",
		NameAbv:     "GOL",
		Line:        "red",
		Coordinates: []float64{53.33589167, -6.313569444},
	},
	&Stop{
		Name:        "Suir Road",
		NameAbv:     "SUI",
		Line:        "red",
		Coordinates: []float64{53.33661667, -6.307211111},
	},
	&Stop{
		Name:        "Rialto",
		NameAbv:     "RIA",
		Line:        "red",
		Coordinates: []float64{53.33790833, -6.297241667},
	},
	&Stop{
		Name:        "Fatima",
		NameAbv:     "FAT",
		Line:        "red",
		Coordinates: []float64{53.33843889, -6.292547222},
	},
	&Stop{
		Name:        "James's",
		NameAbv:     "JAM",
		Line:        "red",
		Coordinates: []float64{53.34194167, -6.293361111},
	},
	&Stop{
		Name:        "Heuston",
		NameAbv:     "HEU",
		Line:        "red",
		Coordinates: []float64{53.34664722, -6.291808333},
	},
	&Stop{
		Name:        "Museum",
		NameAbv:     "MUS",
		Line:        "red",
		Coordinates: []float64{53.34786667, -6.286713889},
	},
	&Stop{
		Name:        "Smithfield",
		NameAbv:     "SMI",
		Line:        "red",
		Coordinates: []float64{53.34713333, -6.277727778},
	},
	&Stop{
		Name:        "Four Courts",
		NameAbv:     "FOU",
		Line:        "red",
		Coordinates: []float64{53.34686389, -6.273436111},
	},
	&Stop{
		Name:        "Hospital",
		NameAbv:     "HOS",
		Line:        "red",
		Coordinates: []float64{53.28936944, -6.37885},
	},
	&Stop{
		Name:        "Jervis",
		NameAbv:     "JER",
		Line:        "red",
		Coordinates: []float64{53.34768611, -6.265333333},
	},
	&Stop{
		Name:        "Abbey Street",
		NameAbv:     "ABB",
		Line:        "red",
		Coordinates: []float64{53.34858889, -6.258172222},
	},
	&Stop{
		Name:        "Busaras",
		NameAbv:     "BUS",
		Line:        "red",
		Coordinates: []float64{53.350075, -6.25145},
	},
	&Stop{
		Name:        "Connolly",
		NameAbv:     "CON",
		Line:        "red",
		Coordinates: []float64{53.35092222, -6.249941667},
	},
	&Stop{
		Name:        "St. Stephen's Green",
		NameAbv:     "STS",
		Line:        "green",
		Coordinates: []float64{53.33907222, -6.261333333},
	},
	&Stop{
		Name:        "Harcourt",
		NameAbv:     "HAR",
		Line:        "green",
		Coordinates: []float64{53.33335833, -6.26265},
	},
	&Stop{
		Name:        "Charlemont",
		NameAbv:     "CHA",
		Line:        "green",
		Coordinates: []float64{53.33066944, -6.258683333},
	},
	&Stop{
		Name:        "Ranelagh",
		NameAbv:     "RAN",
		Line:        "green",
		Coordinates: []float64{53.32643333, -6.256202778},
	},
	&Stop{
		Name:        "Beechwood",
		NameAbv:     "BEE",
		Line:        "green",
		Coordinates: []float64{53.32082222, -6.254652778},
	},
	&Stop{
		Name:        "Cowper",
		NameAbv:     "COW",
		Line:        "green",
		Coordinates: []float64{53.31646667, -6.253447222},
	},
	&Stop{
		Name:        "Cookstown",
		NameAbv:     "COO",
		Line:        "red",
		Coordinates: []float64{53.29350556, -6.384397222},
	},
	&Stop{
		Name:        "Milltown",
		NameAbv:     "MIL",
		Line:        "green",
		Coordinates: []float64{53.30991667, -6.251727778},
	},
	&Stop{
		Name:        "Windy Arbour",
		NameAbv:     "WIN",
		Line:        "green",
		Coordinates: []float64{53.30155833, -6.250708333},
	},
	&Stop{
		Name:        "Dundrum",
		NameAbv:     "DUN",
		Line:        "green",
		Coordinates: []float64{53.29235833, -6.245116667},
	},
	&Stop{
		Name:        "Balally",
		NameAbv:     "BAL",
		Line:        "green",
		Coordinates: []float64{53.28610556, -6.236772222},
	},
	&Stop{
		Name:        "Kilmacud",
		NameAbv:     "KIL",
		Line:        "green",
		Coordinates: []float64{53.28300833, -6.223886111},
	},
	&Stop{
		Name:        "Stillorgan",
		NameAbv:     "STI",
		Line:        "green",
		Coordinates: []float64{53.27931111, -6.209919444},
	},
	&Stop{
		Name:        "Sandyford",
		NameAbv:     "SAN",
		Line:        "green",
		Coordinates: []float64{53.27760278, -6.204677778},
	},
	&Stop{
		Name:        "Central Park",
		NameAbv:     "CPK",
		Line:        "green",
		Coordinates: []float64{53.27015, -6.203763889},
	},
	&Stop{
		Name:        "Glencairn",
		NameAbv:     "GLE",
		Line:        "green",
		Coordinates: []float64{53.26633611, -6.209941667},
	},
	&Stop{
		Name:        "The Gallops",
		NameAbv:     "GAL",
		Line:        "green",
		Coordinates: []float64{53.26116389, -6.206022222},
	},
	&Stop{
		Name:        "Belgard",
		NameAbv:     "BEL",
		Line:        "red",
		Coordinates: []float64{53.29928611, -6.374886111},
	},
	&Stop{
		Name:        "Leopardstown Valley",
		NameAbv:     "LEO",
		Line:        "green",
		Coordinates: []float64{53.25824722, -6.198361111},
	},
	&Stop{
		Name:        "Ballyogan Wood",
		NameAbv:     "GAW",
		Line:        "green",
		Coordinates: []float64{53.255047, -6.184475},
	},
	&Stop{
		Name:        "Carrickmines",
		NameAbv:     "CCK",
		Line:        "green",
		Coordinates: []float64{53.25403333, -6.169908333},
	},
	&Stop{
		Name:        "Laughanstown",
		NameAbv:     "LAU",
		Line:        "green",
		Coordinates: []float64{53.25060556, -6.155005556},
	},
	&Stop{
		Name:        "Cherrywood",
		NameAbv:     "CHE",
		Line:        "green",
		Coordinates: []float64{53.24533333, -6.145852778},
	},
	&Stop{
		Name:        "Bride's Glen",
		NameAbv:     "BRI",
		Line:        "green",
		Coordinates: []float64{53.242075, -6.142886111},
	},
	&Stop{
		Name:        "Fettercairn",
		NameAbv:     "FET",
		Line:        "red",
		Coordinates: []float64{53.29351885, -6.395553517},
	},
	&Stop{
		Name:        "Kingswood",
		NameAbv:     "KIN",
		Line:        "red",
		Coordinates: []float64{53.30369444, -6.36525},
	},
	&Stop{
		Name:        "Cheeverstown",
		NameAbv:     "CVN",
		Line:        "red",
		Coordinates: []float64{53.29098242, -6.4068485},
	},
	&Stop{
		Name:        "Citywest Campus",
		NameAbv:     "CIT",
		Line:        "red",
		Coordinates: []float64{53.28783255, -6.418914583},
	},
	&Stop{
		Name:        "Fortunestown",
		NameAbv:     "FOR",
		Line:        "red",
		Coordinates: []float64{53.28425063, -6.42460165},
	},
	&Stop{
		Name:        "Saggart",
		NameAbv:     "SAG",
		Line:        "red",
		Coordinates: []float64{53.28467885, -6.43776255},
	},
	&Stop{
		Name:        "George's Dock",
		NameAbv:     "GDK",
		Line:        "red",
		Coordinates: []float64{53.349528, -6.247575},
	},
	&Stop{
		Name:        "Mayor Square - NCI",
		NameAbv:     "MYS",
		Line:        "red",
		Coordinates: []float64{53.34924722, -6.243394444},
	},
	&Stop{
		Name:        "Spencer Dock",
		NameAbv:     "SDK",
		Line:        "red",
		Coordinates: []float64{53.34882222, -6.237147222},
	},
	&Stop{
		Name:        "The Point",
		NameAbv:     "TPT",
		Line:        "red",
		Coordinates: []float64{53.34835, -6.229258333},
	},
	&Stop{
		Name:        "Dawson",
		NameAbv:     "DAW",
		Line:        "green",
		Coordinates: []float64{53.34216927, -6.25796366},
	},
	&Stop{
		Name:        "Red Cow",
		NameAbv:     "RED",
		Line:        "red",
		Coordinates: []float64{53.31683333, -6.369872222},
	},
	&Stop{
		Name:        "Trinity",
		NameAbv:     "TRI",
		Line:        "green",
		Coordinates: []float64{53.34528172, -6.25825674},
	},
	&Stop{
		Name:        "Westmoreland",
		NameAbv:     "WES",
		Line:        "green",
		Coordinates: []float64{53.34635626, -6.25897083},
	},
	&Stop{
		Name:        "Marlborough",
		NameAbv:     "MAR",
		Line:        "green",
		Coordinates: []float64{53.34924219, -6.25773649},
	},
	&Stop{
		Name:        "O'Connell - GPO",
		NameAbv:     "OGP",
		Line:        "green",
		Coordinates: []float64{53.34884509, -6.25988168},
	},
	&Stop{
		Name:        "O'Connell - Upper",
		NameAbv:     "OUP",
		Line:        "green",
		Coordinates: []float64{53.35161218, -6.26102943},
	},
	&Stop{
		Name:        "Parnell",
		NameAbv:     "PAR",
		Line:        "green",
		Coordinates: []float64{53.35310356, -6.26050153},
	},
	&Stop{
		Name:        "Dominick",
		NameAbv:     "DOM",
		Line:        "green",
		Coordinates: []float64{53.3513492, -6.26554234},
	},
	&Stop{
		Name:        "Broadstone - DIT",
		NameAbv:     "BRD",
		Line:        "green",
		Coordinates: []float64{53.35407442, -6.27375642},
	},
	&Stop{
		Name:        "Grangegorman",
		NameAbv:     "GRA",
		Line:        "green",
		Coordinates: []float64{53.35719205, -6.27733848},
	},
	&Stop{
		Name:        "Phibsborough",
		NameAbv:     "PHI",
		Line:        "green",
		Coordinates: []float64{53.36037646, -6.27888498},
	},
	&Stop{
		Name:        "Kylemore",
		NameAbv:     "KYL",
		Line:        "red",
		Coordinates: []float64{53.32665556, -6.343444444},
	},
	&Stop{
		Name:        "Cabra",
		NameAbv:     "CAB",
		Line:        "green",
		Coordinates: []float64{53.36433419, -6.28196899},
	},
	&Stop{
		Name:        "Broombridge",
		NameAbv:     "BRO",
		Line:        "green",
		Coordinates: []float64{53.37223956, -6.29768465},
	},
	&Stop{
		Name:        "Bluebell",
		NameAbv:     "BLU",
		Line:        "red",
		Coordinates: []float64{53.32929722, -6.333791667},
	},
	&Stop{
		Name:        "Blackhorse",
		NameAbv:     "BLA",
		Line:        "red",
		Coordinates: []float64{53.33425833, -6.327394444},
	},
}