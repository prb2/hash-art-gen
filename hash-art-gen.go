package main

import (
    "crypto/sha256"
    "flag"
    "fmt"
)



func main() {

    var str_in string
    flag.StringVar(&str_in, "seed", "", "provide a short string to use as a seed, otherwise a random seed will be used")
    flag.Parse()

    if len(str_in) == 0 {
        // TODO: make this a random string
        str_in = "helloworld"
    }

    h := sha256.New()
    h.Write([]byte(str_in))

    hash := h.Sum(nil)
    fmt.Printf("Generating random art with seed\n\n\t%s\n\n", str_in)
    fmt.Printf("SHA256 hash\n\n\t%x\n\n", hash)

    //fmt.Printf("b: %b\n", h.Sum(nil))
    //fmt.Printf("v: %v\n", h.Sum(nil))

    augmentation_string := " .o+=*BOX@%&#/^"
    aug_slice := make([]rune, len(augmentation_string))
    for  i, c := range augmentation_string {
        aug_slice[i] = c
    }
    fmt.Printf("Augmentation runes\n\n\t%c\n\n", aug_slice)

    fmt.Println("Gen art\n")
    gen_art_from_hash(hash, aug_slice)
}

func gen_art_from_hash(hash []byte, aug_slice []rune) {

    //fmt.Printf("%b, is type: %T, len: %v\n", h.Sum(nil), h.Sum(nil), len(h.Sum(nil)))

    //as_bytes := []byte(h.Sum(nil))
    //fmt.Printf("%b, is type: %T, len: %v\n", as_bytes, as_bytes, len(as_bytes))

    bits := ""
    for _, n := range hash {
        as_bin := fmt.Sprintf("%08b", n)
        //leadingz := bits.LeadingZeros(uint(n))
        //with_lead := ("0" * leadingz) + as_bin
        //fmt.Println("Orig: ", n, "As bin: ", as_bin, "with leading: ", with_lead)
        //fmt.Println("Orig: ", n, "As bin: ", as_bin)
        bits += as_bin
    }
    //fmt.Println(bits, "len", len(bits))

    grid := [9][17]int{}
    num_rows := len(grid)
    num_cols := len(grid[0])
    //fmt.Printf("Grid: [%v x %v]\n", num_rows, num_cols)

    within_bound := func(i, j int) bool {
        return i >= 0 && i < num_rows && j >= 0 && j < num_cols
    }

    move_up_left := func(i, j int, grid *[9][17]int) (int, int) {
        movement := ""
        if within_bound(i - 1, j) {
            i -= 1
            movement += "up"
        }
        if within_bound(i, j - 1) {
            j -= 1
            movement += "left"
        }
        if len(movement) == 0 {
            movement = "stayed"
        }
        //fmt.Printf("%v to (%v, %v)\n", movement, i, j)
        grid[i][j] += 1
        return i, j
    }

    move_up_right := func(i, j int, grid *[9][17]int) (int, int) {
        movement := ""
        if within_bound(i - 1, j) {
            i -= 1
            movement += "up"
        }
        if within_bound(i, j + 1) {
            j += 1
            movement += "right"
        }
        if len(movement) == 0 {
            movement = "stayed"
        }
        //fmt.Printf("%v to (%v, %v)\n", movement, i, j)
        grid[i][j] += 1
        return i, j
    }

    move_down_left := func(i, j int, grid *[9][17]int) (int, int) {
        movement := ""
        if within_bound(i + 1, j) {
            i += 1
            movement += "down"
        }
        if within_bound(i, j - 1) {
            j -= 1
            movement += "left"
        }
        if len(movement) == 0 {
            movement = "stayed"
        }
        //fmt.Printf("%v to (%v, %v)\n", movement, i, j)
        grid[i][j] += 1
        return i, j
    }

    move_down_right := func(i, j int, grid *[9][17]int) (int, int) {
        movement := ""
        if within_bound(i + 1, j) {
            i += 1
            movement += "down"
        }
        if within_bound(i, j + 1) {
            j += 1
            movement += "right"
        }
        if len(movement) == 0 {
            movement = "stayed"
        }
        //fmt.Printf("%v to (%v, %v)\n", movement, i, j)
        grid[i][j] += 1
        return i, j
    }
    //fmt.Println(within_bound(100, -1))
    //fmt.Println(within_bound(1, 3))
    //fmt.Println(within_bound(1, 17))
    //fmt.Println(within_bound(9, 3))

    x := 4
    y := 9
    for i := 0; i < len(bits); i += 2 {
        move := bits[i:i+2]
        //fmt.Printf("%v\n", move)

        switch move {
            case "00": {
                x, y = move_up_left(x, y, &grid)
            }
            case "01": {
                x, y = move_up_right(x, y, &grid)
            }
            case "10": {
                x, y = move_down_left(x, y, &grid)
            }
            case "11": {
                x, y = move_down_right(x, y, &grid)
            }
        }

    }

    //print_grid(&grid)

    print_grid_runes(&grid, aug_slice)
}


func print_grid(grid *[9][17]int) {
    for i := 0; i < len(grid); i++ {
        fmt.Println(grid[i])
    }
}

func print_grid_runes(grid *[9][17]int, aug_slice []rune) {

    fmt.Println("\t+-------------------+")
    for i := 0; i < len(grid); i++ {
        row := ""
        for j := 0; j < len(grid[i]); j++ {
            row += fmt.Sprintf("%c", aug_slice[grid[i][j]])
        }
        fmt.Println("\t|", row, "|")
    }
    fmt.Println("\t+-------------------+\n\n")
}
