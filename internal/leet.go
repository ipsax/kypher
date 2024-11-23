package internal

import (
	"fmt"
)

type Leet struct{}

// todo make type for leetchars = https://en.wikipedia.org/wiki/Leet#Orthography
func leetIterator() {
	var a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z []string
	a = append(a, "a", "4", "/", `\`, `@`, "/-", "^", "(L", "Д")
	b = append(b, "b")
	c = append(c, "c")
	d = append(d, "d")
	e = append(e, "e")
	f = append(f, "f")
	g = append(g, "g")
	h = append(h, "h")
	i = append(i, "i")
	j = append(j, "j")
	k = append(k, "k")
	l = append(l, "l")
	m = append(m, "m")
	n = append(n, "n")
	o = append(o, "o")
	p = append(p, "p")
	q = append(q, "q")
	r = append(r, "r")
	s = append(s, "s")
	t = append(t, "t")
	u = append(u, "u")
	v = append(v, "v")
	w = append(w, "w")
	x = append(x, "x")
	y = append(y, "y")
	z = append(z, "z")

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

	fmt.Println(leetChars)
}

func (Leet) Encode(str string) string {
	leetIterator()
	return "todo return str->leet cipher"
}

func (Leet) Decode(str string) string {
	return "todo return leet->str msg"
}
