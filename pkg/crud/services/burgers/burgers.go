package burgers

import (
	"context"
	"errors"
	errors2 "github.com/ParvizBoymurodov/web/pkg/crud/errors"
	"github.com/ParvizBoymurodov/web/pkg/crud/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BurgersSvc struct {
	pool *pgxpool.Pool
}

func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	list = make([]models.Burger, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil,errors2.QueryErrors("can't execute pool: ",err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT id, name, price FROM burgers WHERE removed = FALSE")
	if err != nil {
		return nil, errors2.QueryErrors("can't query: ",err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Burger{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, errors2.QueryErrors("can't scan: ",err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *BurgersSvc) Save(model models.Burger) (err error) {
	save, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors2.QueryErrors("can't execute pool: ",err)
	}
	defer save.Release()
	_, err = save.Exec(context.Background(), "INSERT INTO burgers(name,price) VALUES ($1,$2)", model.Name, model.Price)
	if err != nil {
		return errors2.QueryErrors("can't save:  ",err)
	}

	return nil
}

func (service *BurgersSvc) RemoveById(id int) (err error) {
	remove, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors2.QueryErrors("can't execute pool: ",err)
	}
	defer remove.Release()
	_, err = remove.Exec(context.Background(), "UPDATE burgers SET removed = true WHERE id = $1", id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ",err)
	}
	return nil
}
