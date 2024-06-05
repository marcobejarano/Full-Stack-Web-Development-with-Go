package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"

    chapter_01 "fitness.dev/app/gen"
    _ "github.com/lib/pq"
)

func main() {
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
        panic(err)
    }

    // Connectivity check
    if err := db.Ping(); err != nil {
        log.Fatalln("Error from database ping:", err)
    }

    // Create the store
    st := chapter_01.New(db)

    ctx := context.Background()

    _, err = st.CreateUsers(ctx, chapter_01.CreateUsersParams{
        UserName:     "testuser",
        PassWordHash: "hash",
        Name:         "test",
    })

    if err != nil {
        log.Fatalln("Error creating user :", err)
    }

    eid, err := st.CreateExercise(ctx, "Exercise1")

    if err != nil {
        log.Fatalln("Error creating exercise :", err)
    }

    set, err := st.CreateSet(ctx, chapter_01.CreateSetParams{
        ExerciseID: eid,
        Weight:     100,
    })

    if err != nil {
        log.Fatalln("Error updating exercise :", err)
    }

    set, err = st.UpdateSet(ctx, chapter_01.UpdateSetParams{
        ExerciseID: eid,
        SetID:      set.SetID,
        Weight:     2000,
    })

    if err != nil {
        log.Fatalln("Error updating set :", err)
    }

    log.Println("Done!")

    u, _ := st.ListUsers(ctx)

    for _, usr := range u {
        fmt.Printf("Name : %s, ID : %d\n", usr.Name, usr.UserID)
    }
}
