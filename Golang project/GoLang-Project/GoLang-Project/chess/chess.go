package chess

import (
	"math/rand"
	"strconv"
	"time"
)

type Output struct {
	Fen     string `json:"fen"`
	Message string `json:"message"`
}

var charset = [37]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var Pieces []int = []int{'K', 'Q', 'R', 'q', 'N', 'B', 'n', 'k', 'b', 'r', 'p', 'P'}

var elementcode []string = []string{"K", "N2", "P1", "P6", "R2", "P4", "B2", "P5", "r1", "p1", "n2", "p6", "k",
	"p4", "P2", "R1", "b2", "p2", "n1", "P8", "p3", "b1", "P7", "p8", "Q", "q", "P3", "B1", "p5", "N1", "r2", "p7"}

var used []string

var positionOfElement []int

type Pos struct {
	//Character   string
	Element     int
	ElementCode string
	Index       []string
}

/*func (p Pos) getKeyByIndex(index string) bool {
	if contains(p.Index, index) {
		return true
	}
}
*/
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getCodeByIndex(index string) string {
	for _, value := range m {
		if contains(value.Index, index) {
			return value.ElementCode
		}
	}
	return ""
}

func getKeyByIndex(index string) string {
	for key, value := range m {
		if contains(value.Index, index) {
			return key
		}
	}
	return ""
}

var m = make(map[string]Pos) //index of array
func SetPositionOfElements() {

	for i := 0; i < 6; i++ {

		m[strconv.Itoa(i)] = Pos{int(elementcode[i][0]), elementcode[i], []string{"0" + strconv.Itoa(i)}}
	}

	indexOfAlpha := 10 // Starts from A
	elementIndex := 6
	for i := 1; i < 5; i++ { //1-4
		for j := 0; j < 8; j++ {
			if i == 1 && (j == 0 || j == 1) { // ignore for 10 and 11 index
				continue
			}
			if i == 4 && j == 4 { // break loop at 44 index
				break
			}
			m[charset[indexOfAlpha]] = Pos{int(elementcode[elementIndex][0]), elementcode[elementIndex], []string{"" + strconv.Itoa(i) + strconv.Itoa(j)}}
			indexOfAlpha++
			elementIndex++
		}
	}

	indexOfAlpha = 10
	elementIndex = 6
	for i := 4; i < 8; i++ { //4-7
		for j := 0; j < 8; j++ {
			if i == 4 && (j == 0 || j == 1 || j == 2 || j == 3) {
				continue
			}
			if i == 7 && j == 6 {
				break
			}
			currchar := charset[indexOfAlpha]
			m[currchar] = Pos{m[currchar].Element, m[currchar].ElementCode, append(m[currchar].Index, ("" + strconv.Itoa(i) + strconv.Itoa(j)))}
			indexOfAlpha++
		}
	}

	//fmt.Println("map :", m)
}

func Encode(playerChoice string) Output {

	var result Output

	SetPositionOfElements()

	rand.Seed(time.Now().UnixNano())
	used = nil
	for i := 0; i < len(playerChoice); i++ {
		charcode := string(playerChoice[i])
		var index string
		if charcode == "0" || charcode == "1" || charcode == "2" || charcode == "3" || charcode == "4" || charcode == "5" { // for 0-5 no need to randomize
			index = m[charcode].Index[0]
		} else { // out of 2 available indexes pick one
			index = m[charcode].Index[rand.Intn(2)]
		}
		used = append(used, index)
	}
	//fmt.Println(used)

	// generate fen code
	var str string
	var counter int
	str = ""
	var s string
	for i := 0; i < 8; i++ {
		counter = 0
		for j := 0; j < 8; j++ {
			idx := "" + strconv.Itoa(i) + strconv.Itoa(j)
			s = getCodeByIndex(idx)
			if s != "" && contains(used, idx) {
				if counter != 0 {
					str += "" + strconv.Itoa(counter)
				}
				str += "" + string(s[0])
				counter = 0
			} else {
				counter++
			}

		}
		if counter != 0 { // in case counter variable is at last
			str += "" + strconv.Itoa(counter)
		}
		str += "/"
	}

	str = str[0 : len(str)-1]
	result.Fen = str
	return result
}

func Decode(playerChoice string) Output {
	message := ""
	for i := 0; i < len(used); i++ {
		message += getKeyByIndex(used[i])
	}

	var result Output

	result.Fen = playerChoice

	result.Message = message
	return result
}
