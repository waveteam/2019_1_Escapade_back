package game

import (
	"sync"
)

// SendPredicate - returns true if the parcel send to that conn
type SendPredicate func(conn *Connection) bool

// SendToConnections send 'info' to everybody,  whose predicate
// returns true
func SendToConnections(info interface{},
	predicate SendPredicate, groups ...[]*Connection) {

	waitJobs := &sync.WaitGroup{}
	for _, group := range groups {
		for _, connection := range group {
			if predicate(connection) {
				waitJobs.Add(1)
				go connection.sendGroupInformation(info, waitJobs)
			}
		}
	}
	waitJobs.Wait()
}

// AllExceptThat is SendPredicate to SendToConnections
// it will send everybody except selected one and disconnected
func AllExceptThat(me *Connection) func(*Connection) bool {
	return func(conn *Connection) bool {
		return !conn.done() && !me.done() && conn.ID() != me.ID() && conn.IsConnected()
	}
}

// All is SendPredicate to SendToConnections
// it will send everybody, who is connected
func All(conn *Connection) bool {
	return !conn.done() && conn.IsConnected()
}

// Me is SendPredicate to SendToConnections
// it will send only to selected connection
func Me(me *Connection) func(*Connection) bool {
	return func(conn *Connection) bool {
		return !conn.done() && !me.done() && conn.ID() == me.ID() && conn.IsConnected()
	}
}

// All is SendPredicate to SendToConnections
// it will send everybody in room, who is connected
func (room *Room) All(conn *Connection) bool {
	return !conn.done() && conn.Room() == room && conn.IsConnected()
}

// InGame is SendPredicate to SendToConnections
// it will send everybody in room, if game began
func (room *Room) InGame(conn *Connection) bool {
	return !conn.done() && conn.Room() == room && conn.IsConnected() && !conn.Both()
}

// AllExceptThat is SendPredicate to SendToConnections
// it will send everybody in room, except selected one
func (room *Room) AllExceptThat(me *Connection) func(*Connection) bool {
	return func(conn *Connection) bool {
		return !conn.done() && !me.done() && conn.Room() == room && conn != me && conn.IsConnected()
	}
}
