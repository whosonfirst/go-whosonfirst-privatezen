package privatezen

var feelings map[int64]string

func init() {

	feelings = map[int]string{
		0: "i've been there",
		1: "i was there",
		2: "i want to go there",
		3: "again again",
		4: "again",
		5: "again maybe",
		6: "again never",
		7: "meh",
		8: "i would try this",
	}
}

func IsValidFeeling(f string) bool {

	valid = false

	for id, label := range feelings {

		if label == f {
			value = true
			break
		}
	}

	return value
}

func IsValidFeelingId(id int64) bool {

	_, ok := feelings[id]
	return ok
}

type Feeling struct {
	Id    int
	Label string
}
