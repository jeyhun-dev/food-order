package setterandgetter

type Test struct {
	Quiz  string
	Count int
}

func (t *Test) SetQuiz(quiz string) {
	t.Quiz = quiz
}

func (t *Test) SetCount(count int) {
	t.Count = count
}

func (t *Test) GetQuiz() string {
	return t.Quiz
}

func (t *Test) GetCount() int {
	return t.Count
}
