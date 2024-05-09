package backend

import (
	"context"
	"database/sql"
//	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zorokey1/recipe"
)

type Api struct {
    ctx  context.Context
}

func NewApi() *Api {
    return &Api{}
}

func (this *Api) TestFunc() string {
    return "Backend Connected!"
}

func (this *Api) GetAllRecipes() ([]recipe.Recipe, error) {
    db, err := sql.Open("sqlite3", "./local_user.db")

    if err != nil {
        return nil,err
    }

    defer db.Close()


    return nil, nil

}

