package politics

// Question represtens a question to determine the policital orientation
type Question struct {
	text             string
	economyEffect    int8
	diplomacyEffect  int8
	governmentEffect int8
	societyEffect    int8
}
