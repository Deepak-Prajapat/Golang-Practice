package main

type Game struct {
	Name  string
	Class string
}

/* func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout

	cmd.Run()
}
func main() {
	clearTerminal()
	var gameArray []*Game = []*Game{}
	fmt.Println("gameArray", gameArray)
	gameArray = append(gameArray, &Game{Name:"Raghav", Class: "BCA"})
	gameArray = append(gameArray, &Game{Name:"Deepak", Class: "BCA"})
	fmt.Println("gameArray", gameArray)
}
*/
