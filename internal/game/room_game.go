package game

import (
	"fmt"

	"github.com/go-park-mail-ru/2019_1_Escapade/internal/models"
	re "github.com/go-park-mail-ru/2019_1_Escapade/internal/return_errors"
)

// Winner determine who won the game
func (room *Room) Winner() (idWin int) {
	max := 0

	for id, player := range room.Players.Players {
		if player.Points > max {
			max = player.Points
			idWin = id
		}
	}
	return
}

// flagFound is called, when somebody find cell flag
func (room *Room) flagFound(founder Connection, found *Cell) bool {
	thatID := found.Value - CellIncrement
	if thatID == founder.ID() {
		return false
	}
	room.Players.Players[founder.Index].Points += 100000
	fmt.Println("start search!")
	for _, conn := range room.Players.Connections {
		fmt.Println("compare:", thatID, conn.ID())
		if thatID == conn.ID() {
			return room.kill(conn, ActionFlagLost)
		}
	}
	fmt.Println("finish search!")
	return false
}

// isAlive check if connection is player and he is not died
func (room *Room) isAlive(conn *Connection) bool {
	return conn.Index >= 0 && !room.Players.Players[conn.Index].Finished
}

// setFinished increment amount of killed
func (room *Room) setFinished(conn *Connection) {
	room.Players.Players[conn.Index].Finished = true
	room.Players.Players[conn.Index].Died = true
	room.killed++
}

// kill make user die and check for finish battle
func (room *Room) kill(conn *Connection, action int) bool {
	// cause all in pointers
	if room.Status != StatusRunning {
		return false
	}
	if room.isAlive(conn) {
		room.Field.setCellFlagTaken(&room.Players.Flags[conn.Index])

		room.setFinished(conn)
		if room.Players.Capacity <= room.killed+1 {
			room.finishGame(true)
		}
		pa := *room.addAction(conn.ID(), action)
		room.sendAction(pa, room.All)
	}
	return room.Players.Capacity <= room.killed+1
}

// GiveUp kill connection, that call it
func (room *Room) GiveUp(conn *Connection) {
	room.kill(conn, ActionGiveUp)
}

// flagExists find players with such flag. This - flag owner
func (room *Room) flagExists(cell Cell, this *Connection) (found bool, conn Connection) {
	var player int
	for index, flag := range room.Players.Flags {
		if (flag.X == cell.X) && (flag.Y == cell.Y) {
			if this == nil || index != this.Index {
				found = true
				player = index
			}
			break
		}
	}
	if !found {
		return
	}
	for _, connection := range room.Players.Connections {
		if connection.Index == player {
			conn = *connection
			break
		}
	}
	return
}

func (room *Room) setFlagCoordinates(conn Connection, cell Cell) {
	room.Players.Flags[conn.Index].X = cell.X
	room.Players.Flags[conn.Index].Y = cell.Y
}

func (room *Room) setAndSendNewCell(conn Connection) {
	found := true
	// create until it become unique
	var cell Cell
	for found {
		cell = room.Field.CreateRandomFlag(conn.ID())
		found, _ = room.flagExists(cell, nil)
	}
	room.setFlagCoordinates(conn, cell)
	response := models.RandomFlagSet(cell)
	conn.SendInformation(response)
}

// setFlag handle user wanna set flag
func (room *Room) setFlag(conn *Connection, cell *Cell) bool {
	// if user try set flag after game launch
	if room.Status != StatusFlagPlacing {
		response := models.FailFlagSet(cell, re.ErrorBattleAlreadyBegan())
		conn.SendInformation(response)
		return false
	}

	if !room.Field.IsInside(cell) {
		response := models.FailFlagSet(cell, re.ErrorCellOutside())
		conn.SendInformation(response)
		return false
	}

	if !room.isAlive(conn) {
		response := models.FailFlagSet(cell, re.ErrorPlayerFinished())
		conn.SendInformation(response)
		return false
	}

	if found, prevConn := room.flagExists(*cell, conn); found {
		room.setAndSendNewCell(*conn)
		room.setAndSendNewCell(prevConn)
		return true
	}

	room.setFlagCoordinates(*conn, *cell)
	return true
}

// setFlags set players flags to field
// call it if game has already begun
func (room *Room) setFlags() {
	for _, cell := range room.Players.Flags {
		room.Field.SetFlag(cell.X, cell.Y, cell.PlayerID)
	}
}

// fillField set flags and mines
func (room *Room) fillField() {
	fmt.Println("fillField", room.Field.Height, room.Field.Width, len(room.Field.Matrix))

	room.setFlags()
	room.Field.SetMines()

}
