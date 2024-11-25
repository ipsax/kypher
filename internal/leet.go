package internal

import (
	"math/rand"
)

type Leet struct{}

// todo make type for leetchars = https://en.wikipedia.org/wiki/Leet#Orthography
func leetMap() map[string][]string {
	var a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z []string
	a = append(a, "a", "4", "/", `\`, `@`, "/-", "^", "(L", "Д")
	b = append(b, "b", "I3", "8", "13", "|3", "ß", "!3", "(3", "/3", ")3", "|-]", "j3")
	c = append(c, "c", "[", "¢", "<", "(", "©")
	d = append(d, "d", ")", "|)", "(|", "[)", "I>", "|", ">", "?", "T)", "I7", "cl", "|}", "|]")
	e = append(e, "e", "3", "&", "£", "€", "[-", "|", "=-")
	f = append(f, "f", "|", "=", "ƒ", "#", "ph", "/", "=", "v")
	g = append(g, "g", "6", "&", "(_+", "9", "C-", "gee", "(?,", "[,", "{,", "<-", "(.")
	h = append(h, "h", "#", "/-/", "\\-\\", "[-]", "]-[", ")-(", "(-)", ":-:", "|~|", "|-|", "]~[", "}{", "!-!", "1-1", "\\-/", "I+I", "?")
	i = append(i, "i", "1", "|", "][", "!", "eye", "3y3")
	j = append(j, "j", ",_|", "_|", "._|", "._]", "_]", ",_]", "]")
	k = append(k, "k", ">|", "|<", "1<", "|c", "|", "(7<")
	l = append(l, "l", "1", "7", "2", "£", "|", "_")
	m = append(m, "m", `\/`, `V\\`, `[V]`, `\\/|`, "^^", `<\\/`, `{V}`, `(v)`, `(V)`, `\|`, `]\\/`, `nn`, `11`)
	n = append(n, "n", "^/", `\|`, `\/`, `[\]`, `<\>`, `{\}`, "/V", "^", "ท", "И")
	o = append(o, "o", "0", "()", "oh", "[]", "p", "<>", "Ø")
	p = append(p, "p", "|*", "|o", "|º", "|^", "|>", `|"`, "9", `[]D`, "|7")
	q = append(q, "q", "(_,)", "()_", "2", "0_", "&", "¶", "⁋", "℗")
	r = append(r, "r", "I2", "9", "|`", "|~", "|?", "/2", "|^", "lz", "7", "2", "12", "®", "[z", "Я", "3", "4", "|2", ".-")
	s = append(s, "s", "5", "$", "z", "§", "ehs", "es", "2")
	t = append(t, "t", "7", "+", "-|-", "']", "['", "†", "«|»", "~|", "~")
	u = append(u, "u", "(_)", "v", "L|", "บ")
	v = append(v, "v", `\/`, "|/", `\|`)
	w = append(w, "w", `\/\/`, "vv", `\N`, `'//`, `\\'`, `\^/`, "(n)", `\V/`, `\X/`, `\|/`, `\_:_/`, "uu", "2u", "พ", "₩", "ω", `\\//`, `\\//`)
	x = append(x, "x", "><", "}{", "ecks", "×", "}{", ")(", "][")
	y = append(y, "y", "j", "`/", "¥", "`|΄")
	z = append(z, "z", "%", "s", "7_", "2", ">_", "~/")

	leetChars := make(map[string][]string)
	leetChars["a"] = a
	leetChars["b"] = b
	leetChars["c"] = c
	leetChars["d"] = d
	leetChars["e"] = e
	leetChars["f"] = f
	leetChars["g"] = g
	leetChars["h"] = h
	leetChars["i"] = i
	leetChars["j"] = j
	leetChars["k"] = k
	leetChars["l"] = l
	leetChars["m"] = m
	leetChars["n"] = n
	leetChars["o"] = o
	leetChars["p"] = p
	leetChars["q"] = q
	leetChars["r"] = r
	leetChars["s"] = s
	leetChars["t"] = t
	leetChars["u"] = u
	leetChars["v"] = v
	leetChars["w"] = w
	leetChars["x"] = x
	leetChars["y"] = y
	leetChars["z"] = z

	return leetChars
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func (Leet) Encode(str string) string {
	encoded := ""
	leetChars := leetMap()

	for _, c := range str {
		// todo ignore chars that arent a-z
		chars, exists := leetChars[string(c)]
		if exists {
			selectedLeet := chars[randRange(0, len(chars))]
			encoded += selectedLeet
		}
	}
	return encoded
}

func (Leet) Decode(str string) string {
	return "todo return leet->str msg"
}
