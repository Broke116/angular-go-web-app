package mongo

import mgo "gopkg.in/mgo.v2"

// Session is used to store mongodb session
type Session struct {
	session *mgo.Session
}

// NewSession is used for establishing a new session then returning an instance for Session struct
func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Session{session}, err
}

// Copy is a method of session struct. It returns a copy of the mongodb session.
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

// GetCollection is a method of session struct returning a collection as its return parameter.
func (s *Session) GetCollection(database string, collection string) *mgo.Collection {
	return s.session.DB(database).C(collection)
}

// DropDatabase is a method used for dropping the database
func (s *Session) DropDatabase(database string) error {
	if s.session != nil {
		return s.session.DB(database).DropDatabase()
	}
	return nil
}

// Close is a method used for closing the session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
