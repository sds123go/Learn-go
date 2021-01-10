package tempconv

func S2H(s Sheshidu) HuaShidu {
	return HuaShidu(s*9/5 + 32)
}
func H2S(h HuaShidu) Sheshidu {

	return Sheshidu((h - 32) * 5 / 9)
}
