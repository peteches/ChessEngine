package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

// var STARTINGFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func handleUci(ctx context.Context) <-chan string {
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

func handlePosition(ctx context.Context, args []string) {
}

func processInput(ctx context.Context, cmdChan <-chan string) <-chan string {
	outChan := make(chan string)

	go func() {
		defer close(outChan)

		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			case cmd := <-cmdChan:
				{
					words := strings.Split(cmd, " ")
					switch words[0] {
					case "uci":
						{
							go func() {
								o := handleUci(ctx)
								for x := range o {
									outChan <- x
								}
							}()
						}
					case "position":
						{
							go handlePosition(ctx, words[1:])
						}
					default:
						{
							outChan <- fmt.Sprintf("Received unknown CMD: %s\n", cmd)
						}
					}
				}
			}
		}
	}()

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
	ctx := context.TODO()
	cmdChan := scanForCommands(ctx, os.Stdin)
	toGuiChan := processInput(ctx, cmdChan)

	for x := range toGuiChan {
		fmt.Fprintln(os.Stdout, x)
	}
}
