package main

type landUnit struct {
	quantity int
	fire     int
	shock    int
}

type army struct {
	infantry  []landUnit
	cavalry   []landUnit
	artillery []landUnit
}

/*
func main() {
	your_army := army{
		infantry: []landUnit{
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
		},
	}
	villain_army := army{
		infantry: []landUnit{
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
			{1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1}, {1000, 1, 1},
		},
	}
	fmt.Println(your_army)
	fmt.Println(villain_army)

}
*/
