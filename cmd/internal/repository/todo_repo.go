package repository

import (
	"context"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/config/db"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTodo(ctx context.Context, database *db.Database, todo *domain.Todo) (*domain.Todo, error) {
	insertResult, err := database.Collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func GetAllTodos(ctx context.Context, database *db.Database) ([]domain.Todo, error) {
	var todos []domain.Todo

	cursor, err := database.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo domain.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func CompleteTodo(ctx context.Context, database *db.Database, objectId primitive.ObjectID) error {
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err := database.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(ctx context.Context, database *db.Database, objectId primitive.ObjectID) error {
	filter := bson.M{"_id": objectId}
	_, err := database.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
