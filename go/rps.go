package kata

var (
	o1 string = "Player 1 won!"
	o2 string = "Player 2 won!"
)

func Rps(p1, p2 string) string {
	if p1 == "scissors" && p2 == "paper" {
		return o1
	}
	if p1 == "scissors" && p2 == "rock" {
		return o2
	}
	if p1 == "paper" && p2 == "rock" {
		return o1
	}
	if p1 == "paper" && p2 == "scissors" {
		return o2
	}
	if p1 == "rock" && p2 == "scissors" {
		return o1
	}
	if p1 == "rock" && p2 == "paper" {
		return o2
	}
	return "Draw!"
}
