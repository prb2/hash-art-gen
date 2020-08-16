package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/glog"
)

func main() {
	var (
		str_in string
		border bool
		kind   int
	)
	defer glog.Flush()
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&str_in, "seed", "",
		"provide a short string to use as a seed, otherwise a random seed will be used")
	flag.BoolVar(&border, "border", false, "draw border")
	flag.IntVar(&kind, "kind", 0, "0: ssh chars, 1: symbols, 2: arrows")
	flag.Parse()

	var bytes_in []byte
	if len(str_in) > 0 {
		glog.V(1).Infof("Using input: %s", str_in)
		bytes_in = []byte(str_in)
	} else {
		bytes_in = make([]byte, 100)
		_, err := rand.Read(bytes_in)
		if err != nil {
			glog.Error(err)
			return
		}
		glog.V(1).Infof("Using rand bytes: %v", bytes_in)
	}

	h := sha256.New()
	h.Write(bytes_in)

	hash := h.Sum(nil)
	glog.V(1).Infof("Generating random art with seed\n\n\t%s\n\n", str_in)
	glog.V(1).Infof("SHA256 hash\n\n\t%x\n\n", hash)

	glog.V(3).Infof("b: %b\n", h.Sum(nil))
	glog.V(3).Infof("v: %v\n", h.Sum(nil))

	aug_slice := get_aug(augType(kind))
	glog.V(1).Infof("Augmentation runes\n\n\t%c\n\n", aug_slice)

	glog.V(1).Info("Gen art\n")
	grid := gen_art_from_hash(hash)
	print_grid_runes(&grid, aug_slice, border)
}

type augType int

const (
	ssh = iota
	symbols
	arrows
)

func get_aug(kind augType) []rune {
	var aug string
	switch kind {
	case symbols:
		aug = "TODO"
	case arrows:
		aug = " ↶↠↻↙↫⇝←↯↝⇗⇢⇁⇀⇶"
	case ssh:
		fallthrough
	default:
		aug = " .o+=*BOX@%&#/^"
	}

	aug_slice := make([]rune, len(aug))
	for i, c := range aug {
		aug_slice[i] = c
	}
	return aug_slice
}

func gen_art_from_hash(hash []byte) [9][17]int {
	bits := ""
	for _, n := range hash {
		as_bin := fmt.Sprintf("%08b", n)
		bits += as_bin
	}
	glog.V(3).Info(bits, "len", len(bits))

	grid := [9][17]int{}
	num_rows := len(grid)
	num_cols := len(grid[0])
	glog.V(3).Infof("Grid: [%v x %v]\n", num_rows, num_cols)

	within_bound := func(i, j int) bool {
		return i >= 0 && i < num_rows && j >= 0 && j < num_cols
	}

	move_up_left := func(i, j int, grid *[9][17]int) (int, int) {
		movement := ""
		if within_bound(i-1, j) {
			i -= 1
			movement += "up"
		}
		if within_bound(i, j-1) {
			j -= 1
			movement += "left"
		}
		if len(movement) == 0 {
			movement = "stayed"
		}
		glog.V(3).Infof("%v to (%v, %v)\n", movement, i, j)
		grid[i][j] += 1
		return i, j
	}

	move_up_right := func(i, j int, grid *[9][17]int) (int, int) {
		movement := ""
		if within_bound(i-1, j) {
			i -= 1
			movement += "up"
		}
		if within_bound(i, j+1) {
			j += 1
			movement += "right"
		}
		if len(movement) == 0 {
			movement = "stayed"
		}
		glog.V(3).Infof("%v to (%v, %v)\n", movement, i, j)
		grid[i][j] += 1
		return i, j
	}

	move_down_left := func(i, j int, grid *[9][17]int) (int, int) {
		movement := ""
		if within_bound(i+1, j) {
			i += 1
			movement += "down"
		}
		if within_bound(i, j-1) {
			j -= 1
			movement += "left"
		}
		if len(movement) == 0 {
			movement = "stayed"
		}
		glog.V(3).Infof("%v to (%v, %v)\n", movement, i, j)
		grid[i][j] += 1
		return i, j
	}

	move_down_right := func(i, j int, grid *[9][17]int) (int, int) {
		movement := ""
		if within_bound(i+1, j) {
			i += 1
			movement += "down"
		}
		if within_bound(i, j+1) {
			j += 1
			movement += "right"
		}
		if len(movement) == 0 {
			movement = "stayed"
		}
		glog.V(3).Infof("%v to (%v, %v)\n", movement, i, j)
		grid[i][j] += 1
		return i, j
	}

	x := 4
	y := 9
	for i := 0; i < len(bits); i += 2 {
		move := bits[i : i+2]
		glog.V(3).Infof("%v\n", move)

		switch move {
		case "00":
			{
				x, y = move_up_left(x, y, &grid)
			}
		case "01":
			{
				x, y = move_up_right(x, y, &grid)
			}
		case "10":
			{
				x, y = move_down_left(x, y, &grid)
			}
		case "11":
			{
				x, y = move_down_right(x, y, &grid)
			}
		}

	}
	//print_grid(&grid)
	return grid
}

func print_grid(grid *[9][17]int) {
	for i := 0; i < len(grid); i++ {
		glog.V(1).Info(grid[i])
	}
}

func print_grid_runes(grid *[9][17]int, aug_slice []rune, border bool) {
	fmt.Println()
	if border {
		fmt.Println("+-------------------+")
	}
	for i := 0; i < len(grid); i++ {
		row := ""
		for j := 0; j < len(grid[i]); j++ {
			row += fmt.Sprintf("%c", aug_slice[grid[i][j]])
		}
		if border {
			fmt.Println("|", row, "|")
		} else {
			fmt.Println(row)
		}
	}
	if border {
		fmt.Println("+-------------------+")
	}
	fmt.Println()
}
