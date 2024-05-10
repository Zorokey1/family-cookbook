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

func (this *Api) TestFunc() (string, string) {
    return "Backend Connected!","will this go through?"
}

func (this *Api) GetAllRecipes() ([]recipe.Recipe, error) {
    db, err := sql.Open("sqlite3", "./local_user.db")

    if err != nil {
        return nil,err
    }

    defer db.Close()


    return nil, nil
}

func (this *Api) GetRecipe(id int) (recipe.Recipe, error) {
    return recipe.Recipe{}, nil
}

func (this *Api) SaveRecipe(recipe recipe.Recipe) error {
    return nil
}

func (this *Api) SwapIngredients(r recipe.Recipe, indexArray []int) (recipe.Recipe, error) {
    err := r.SwapIngredients(indexArray)
    return r,err 
}

func (this *Api) DeleteRecipe(recipe recipe.Recipe) error {
    return nil
}
