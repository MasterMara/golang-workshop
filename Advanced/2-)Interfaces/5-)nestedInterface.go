package composition

type student struct {
	name        string
	midtermNote int
	finalNote   int
}

type findAverage interface {
	findAverage(midtermNote, finalNote int) float64
}

type findMinimumFinalNote interface {
	findMinimumFinalNote(students []student) int
}

// 1.WAY Calculation interface imports two interface
type calculations interface {
	findAverage
	findMinimumFinalNote
}

func (*student) findAverage(midtermNote, finalNote int) float64 {

	return ((float64(midtermNote) * 0.4) + (float64(finalNote) * 0.6))
}

func (*student) findMinimumFinalNote(students []student) int {

	minimumNote := students[0].finalNote

	for _, student := range students {
		if student.finalNote < minimumNote {
			minimumNote = student.finalNote
		}
	}

	return minimumNote
}

func main() {

	mustafa := student{name: "Mustafa", midtermNote: 70, finalNote: 90}
	kadir := student{name: "Kadir", midtermNote: 40, finalNote: 80}

	students := make([]student, 0, 2)

	students = append(students, mustafa)
	students = append(students, kadir)

}

/*Some Notes By Mustafa KARACABEY:
1-) A type can implement more than one interface.
2-) Investigate this purpose usage as a example

*/
