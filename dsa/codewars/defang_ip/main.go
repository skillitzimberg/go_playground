package main

func isValidIP(ip string) bool {
	return true
}

func split(ip string) []string {
	bs := []byte(ip)
	for i, v := range bs {
		for v != 46 {
			sub := []byte{}
			sub = append(sub, v)
		}
	}
	return []string{""}
}
