package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {}

// EventDispatcherTestSuite é uma estrutura que representa um conjunto de testes para o EventDispatcher.
type EventDispatcherTestSuite struct {
	suite.Suite
	// Campos compartilhados entre os testes
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

// Antes de executar cada teste (TestEventDispatcher_Register, por exemplo),
// o suite.Run executa automaticamente o método SetupTest(), se ele existir.
// Isso serve para preparar o ambiente antes de cada teste individual.
func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{Name: "TestEvent", Payload: "TestPayload"}
	suite.event2 = TestEvent{Name: "TestEvent2", Payload: "TestPayload2"}
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.eventDispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have two handlers registered")

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0], "First handler should be the same")
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1], "Second handler should be the same")
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_AlreadyRegistered() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err, "Should return an error when trying to register the same handler again")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should still have one handler registered")
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have two handlers registered")

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]), "Should have three handlers registered")

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers), "Should have no handlers registered for the first event")
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have two handlers registered")

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler), "Should return true for existing handler")
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2), "Should return true for existing handler")
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3), "Should return false for non-existing handler")
	assert.False(suite.T(), suite.eventDispatcher.Has("TestEvent2", &suite.handler3), "Should return false for non-existing event")
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Register(suite.event.GetName(), eh2)

	suite.eventDispatcher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have two handlers registered")

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err, "Should not return an error when registering a new handler")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]), "Should have three handlers registered")

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have one handler registered")
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0], "First handler should be the same")

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]), "Should have no handlers registered for the first event")

	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]), "Should have no handlers registered for the second event")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
