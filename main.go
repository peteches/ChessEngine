package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uciInit(w io.Writer) {
	fmt.Fprintln(w, "id name PetechesChessBot 0.0")
	fmt.Fprintln(w, "id author Pete 'Peteches' McCabe")
	fmt.Fprintln(w, "uciok")
}

func process_input(w io.Writer, cmd string) {
	switch cmd {
	case "uci":
		{
			uciInit(w)
		}
	default:
		{
			fmt.Fprintf(w, "Recieved unknown CMD: %s\n", cmd)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()
		process_input(os.Stdout, cmd)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading Stdin: ", err)
	}
}
