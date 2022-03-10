package poker_test

import (
	"bytes"
	"fmt"
	"poker"
	"strings"
	"testing"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		cases := []struct {
			winner string
		}{
			{winner: "Chris"},
			{winner: "Cleo"},
		}

		for _, c := range cases {
			in := strings.NewReader(fmt.Sprintf("7\n%s wins\n", c.winner))
			game := &poker.GameSpy{}

			cli := poker.NewCLI(in, dummyStdOut, game)
			cli.PlayPoker()

			if game.FinishedWith != c.winner {
				t.Errorf("wanted Finish called with %q but got %q", c.winner, game.FinishedWith)
			}
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertGameNotStarted(t, game)
		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}
