package letters

import "math/rand"

type Letter struct {
	char      byte
	frequency int
}

var Vowels = []Letter{
	{'e', 11},
	{'a', 8},
	{'i', 7},
	{'o', 7},
	{'u', 4},
}

var Consonants = []Letter{
	{'r', 8},
	{'t', 7},
	{'n', 7},
	{'s', 6},
	{'l', 5},
	{'c', 5},
	{'d', 3},
	{'p', 3},
	{'m', 3},
	{'h', 3},
	{'g', 2},
	{'b', 2},
	{'f', 2},
	{'y', 2},
	{'w', 1},
	{'k', 1},
	{'v', 1},
	{'x', 1},
	{'z', 1},
	{'j', 1},
	{'q', 1},
}

var VowelFrequency int
var ConsonantFrequency int

func init() {
	for _, letter := range Vowels {
		VowelFrequency += letter.frequency
	}
	for _, letter := range Consonants {
		ConsonantFrequency += letter.frequency
	}
}

func SelectLetter(letters []Letter, totalFrequency int) byte {
	r := rand.Intn(totalFrequency)
	for _, letter := range letters {
		if r < letter.frequency {
			return letter.char
		}
		r -= letter.frequency
	}

	return letters[len(letters)-1].char // Fallback
}
