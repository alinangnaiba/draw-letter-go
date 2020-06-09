package letters

import "fmt"

// Letters define an interface
type Letters interface {
	Draw()
}

// O is a definition of letter O
type O struct {
	Size int
}

// X is a definition of letter X
type X struct {
	Size int
}

// Y is a definition of letter Y
type Y struct {
	Size int
}

// Z is a definition of letter Z
type Z struct {
	Size int
}

// Draw is an implementation of Letters interface
func (o O) Draw() {
	for row := 0; row < o.Size; row++ {
		for col := 0; col < o.Size; col++ {
			if row == 0 || row == o.Size-1 {
				if col > 0 && col < o.Size-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			} else {
				if col == 0 || col == o.Size-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

//Draw is an implementation of Letters interface
func (x X) Draw() {
	for row := 0; row < x.Size; row++ {
		maxcol := x.Size - 1
		for col := 0; col <= maxcol; col++ {
			if col == row || col == maxcol-row {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//Draw is an implementation of Letters interface
func (y Y) Draw() {
	mid := y.Size / 2
	for row := 0; row < y.Size; row++ {
		maxcol := y.Size - 1
		for col := 0; col <= maxcol; col++ {
			if row <= mid {
				if col == row || col == maxcol-row {
					fmt.Print("*")
					continue
				}
			}
			if row > mid && col == mid {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Draw is an implementation of Letters interface
func (z Z) Draw() {
	len := z.Size - 1
	for row := 0; row < z.Size; row++ {
		for col := 0; col <= len; col++ {
			if row == 0 || row == len {
				fmt.Print("*")
				continue
			}

			if col == (len - row) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
