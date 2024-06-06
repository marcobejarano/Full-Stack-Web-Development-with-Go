package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"time"

	chapter_01 "fitness.dev/app/gen"
	"fitness.dev/app/logger"
	_ "github.com/lib/pq"
)

func main() {
	l := flag.Bool("local", false, "true - send to stdout, false - send to logging server")
	flag.Parse()

	logger.SetLoggingOutput(*l)

	logger.Logger.Debugf("Application logging to stdout = %v", *l)
	logger.Logger.Info("Starting the application...")

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "Postgres123"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	// Open the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		logger.Logger.Errorf("Error opening database : %s", err.Error())
	}

	// Connectivity check
	if err := db.Ping(); err != nil {
		logger.Logger.Errorf("Error from database ping: %s", err.Error())
	}

	logger.Logger.Info("Database connection fine")

	// Create the store
	st := chapter_01.New(db)

	ctx := context.Background()

	chuser, err := st.CreateUsers(ctx, chapter_01.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})

	if err != nil {
		logger.Logger.Fatal("Error creating user")
	}
	logger.Logger.Info("Success - user creation")

	eid, err := st.CreateExercise(ctx, "Exercise1")

	if err != nil {
		logger.Logger.Errorf("Error creating exercise")
	}
	logger.Logger.Info("Success - exercise creation")

	sid, err := st.UpsertSet(ctx, chapter_01.UpsertSetParams{
		ExerciseID: eid,
		Weight:     100,
	})

	if err != nil {
		logger.Logger.Errorf("Error updating sets")
	}

	_, err = st.UpsertWorkout(ctx, chapter_01.UpsertWorkoutParams{
		UserID:    chuser.UserID,
		SetID:     sid,
		StartDate: time.Time{},
	})

	if err != nil {
		logger.Logger.Errorf("Error updating workouts")
	}
	logger.Logger.Info("Success - updating workout")

	logger.Logger.Info("Application complete")

	// sentry implement something similar
	// https://github.com/getsentry/sentry-go/blob/master/example/basic/main.go#L50
	defer time.Sleep(1 * time.Second)
}
