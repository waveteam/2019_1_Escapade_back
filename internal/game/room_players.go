package game

import (
	"fmt"

	"github.com/go-park-mail-ru/2019_1_Escapade/internal/metrics"
)

// RecoverPlayer call it in lobby.join if player disconnected
func (room *Room) RecoverPlayer(newConn *Connection) {
	if room.done() {
		return
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	// add connection as player
	room.MakePlayer(newConn, true)
	pa := *room.addAction(newConn.ID(), ActionReconnect)
	room.addPlayer(newConn)
	room.sendAction(pa, room.AllExceptThat(newConn))
	//room.greet(newConn, true)

	return
}

// RecoverObserver recover connection as observer
func (room *Room) RecoverObserver(oldConn *Connection, newConn *Connection) {
	if room.done() {
		return
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	go room.MakeObserver(newConn, true)
	pa := *room.addAction(newConn.ID(), ActionReconnect)
	go room.sendAction(pa, room.AllExceptThat(newConn))
	//go room.greet(newConn, false)

	return
}

// observe try to connect user as observer
func (room *Room) addObserver(conn *Connection) bool {
	if room.done() {
		return false
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	if room.lobby.Metrics() {
		metrics.Players.WithLabelValues(room.ID, conn.User.Name).Inc()
	}

	// if we havent a place
	if !room.Observers.EnoughPlace() {
		conn.debug("Room cant execute request ")
		return false
	}
	conn.debug("addObserver")
	room.MakeObserver(conn, true)

	go room.addAction(conn.ID(), ActionConnectAsObserver)
	go room.sendObserverEnter(*conn, room.AllExceptThat(conn))
	room.lobby.sendRoomUpdate(*room, All)

	return true
}

// EnterPlayer handle player try to enter room
func (room *Room) addPlayer(conn *Connection) bool {
	if room.done() {
		return false
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	if room.lobby.Metrics() {
		metrics.Players.WithLabelValues(room.ID, conn.User.Name).Inc()
	}

	// if room have already started
	// if room.Status != StatusPeopleFinding {
	// 	return false
	// }

	conn.debug("Room(" + room.ID + ") wanna connect you")

	// if room hasnt got places
	if !room.Players.EnoughPlace() {
		conn.debug("Room(" + room.ID + ") hasnt any place")
		return false
	}

	room.MakePlayer(conn, true)

	go room.addAction(conn.ID(), ActionConnectAsPlayer)
	go room.sendPlayerEnter(*conn, room.AllExceptThat(conn))
	go room.lobby.sendRoomUpdate(*room, All)

	if !room.Players.EnoughPlace() {
		room.chanStatus <- StatusFlagPlacing
	}

	return true
}

// MakePlayer mark connection as connected as Player
// add to players slice and set flag inRoom true
func (room *Room) MakePlayer(conn *Connection, recover bool) {
	if room.done() {
		return
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	if room.Status != StatusPeopleFinding {
		room.lobby.waiterToPlayer(conn)
		conn.setBoth(false)
	} else {
		conn.setBoth(true)
	}
	room.Players.Add(conn, false)
	room.greet(conn, true)
	if recover {
		room.sendStatus(Me(conn))
	}
	conn.PushToRoom(room)
}

// MakeObserver mark connection as connected as Observer
// add to observers slice and set flag inRoom true
func (room *Room) MakeObserver(conn *Connection, recover bool) {
	if room.done() {
		return
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	if room.Status != StatusPeopleFinding {
		room.lobby.waiterToPlayer(conn)
		conn.setBoth(false)
	} else {
		conn.setBoth(true)
	}
	room.Observers.Add(conn, false)
	room.greet(conn, false)
	if recover {
		room.sendStatus(Me(conn))
	}
	conn.PushToRoom(room)
}

// Search search connection in players and observers of room
func (room *Room) Search(find *Connection) *Connection {
	found, i := room.Players.SearchConnection(find)
	if i >= 0 {
		fmt.Println("player!", found.Disconnected(), found)
		return found
	}
	found, i = room.Observers.SearchByID(find.ID())
	if i >= 0 {
		fmt.Println("observer!", found.Disconnected())
		return found
	}
	return nil
}

// RemoveFromGame control the removal of the connection from the room
func (room *Room) RemoveFromGame(conn *Connection, disconnected bool) (done bool) {
	if room.done() {
		return
	}
	room.wGroup.Add(1)
	defer func() {
		room.wGroup.Done()
	}()

	//fmt.Println("removeDuringGame before len", len(room._Players.Connections))

	i := room.Players.SearchIndexPlayer(conn)
	if i >= 0 {
		if (room.Status == StatusFlagPlacing || room.Status == StatusRunning) && !disconnected {
			fmt.Println("give up", i)
			room.GiveUp(conn)
		}

		done = room.Players.Remove(conn, disconnected)
		if done {
			room.sendPlayerExit(*conn, room.All)
		}
	} else {
		done = room.Observers.Remove(conn, disconnected)
		if done {
			go room.sendObserverExit(*conn, room.All)
		}
	}
	if !done {
		return done
	}
	fmt.Println("removeDuringGame")
	//fmt.Println("removeDuringGame after len", len(room._Players.Connections))
	fmt.Println("removeDuringGame system says", room.Players.Empty())
	if room.Players.Empty() {
		if room.lobby.Metrics() {
			metrics.Rooms.Dec()
		}

		fmt.Println("room.Players.Empty")
		room.Close()
	} else {
		room.lobby.sendRoomUpdate(*room, All)
	}
	fmt.Println("there")
	return done
}
