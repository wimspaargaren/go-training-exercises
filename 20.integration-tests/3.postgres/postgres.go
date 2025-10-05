package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Todo represents a task with a unique ID, description, and completion status.
type Todo struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Task string
	Done bool
}

// BeforeCreate is a GORM hook that is called before a new record is created.
func (u *Todo) BeforeCreate(_ *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// TodoRepository provides methods to interact with the todos table in the database.
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new instance of TodoRepository with the given GORM DB connection.
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// TableName returns the name of the table used by the TodoRepository.
func (t *TodoRepository) TableName() string {
	return "todos"
}

// Create adds a new Todo record to the database.
func (t *TodoRepository) Create(ctx context.Context, todo *Todo) error {
	return t.db.WithContext(ctx).Create(todo).Error
}

// List retrieves all Todo records from the database.
func (t *TodoRepository) List(ctx context.Context) ([]*Todo, error) {
	var todos []*Todo
	if err := t.db.WithContext(ctx).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// GetByID retrieves a Todo record by its unique ID.
func (t *TodoRepository) GetByID(ctx context.Context, id uuid.UUID) (*Todo, error) {
	var todo Todo
	if err := t.db.WithContext(ctx).First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
