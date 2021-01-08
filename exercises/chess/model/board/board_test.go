epackage board

import (
	
	"github.com/Benlevor/4eail40_2020/exercises/chess/model/coord"
	"github.com/Benlevor/4eail40_2020/exercises/chess/model/piece"
	"github.com/Benlevor/4eail40_2020/exercises/chess/model/player"
	"testing"
	"reflect"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

type mockPiece struct {
}

func (r mockPiece) String() string {
	panic("implement me") //TODO: create function
}

func (r mockPiece) Moves() map[coord.ChessCoordinates]bool {
	panic("implement me") //TODO: create function
}
func (r mockPiece) Color() player.Color {
	panic("implement me") //TODO: create function
}


func TestClassic_MovePiece(t *testing.T) {
	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       *Classic
		args    args
		wantErr bool
	}{
		{
			"testmock",
			&Classic{},
			args{mockCoord(0), mockCoord(1)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_MovePiece(t *testing.T) {
	class := Classic{}
	piece := mockPiece{}
	piece1 := mockPiece{}
	coordin := coord.NewCartesian(0,0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = pi
	coordin1 := coord.NewCartesian(7, 0)
	x1, _ := coordin1.Coord(0)
	y1, _ := coordin1.Coord(1)
	class[x1][y1] = piece1

	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"Pas de pièce à partir de",
			class,
			args{from: coord.NewCartesian(3, 5), to: coordin},
			true,
		},
		{
			"Occupé à",
			class,
			args{from: coordin1, to: coordin},
			true,
		},
		{
			"Accepté",
			class,
			args{from: coordin, to: coord.NewCartesian(4, 4)},
			false,
		},
	}
	for _, tt := range tests 
