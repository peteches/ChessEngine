package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

func handleUci() <-chan string {
	infoLines := []string{
		"id name PetechesChessBot 0.0",
		"id author Pete 'Peteches' McCabe",
		"uciok",
	}

	outChan := make(chan string, len(infoLines))
	defer close(outChan)

	for _, line := range infoLines {
		outChan <- fmt.Sprintln(line)
	}

	return outChan
}

func scanForCommands(ctx context.Context, r io.Reader) <-chan string {
	scanner := bufio.NewScanner(r)
	cmdChan := make(chan string)

	go func() {
		defer close(cmdChan)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				{
					goto end
				}
			default:
				{
					cmd := scanner.Text()
					switch cmd {
					case "quit":
						goto end
					default:
						cmdChan <- cmd
					}
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading Stdin: ", err)
			}
		}
	end:
	}()

	return cmdChan
}

func main() {
}
