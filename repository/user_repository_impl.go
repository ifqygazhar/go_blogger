package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_blogger/helper"
	"go_blogger/model/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) SignUp(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	SQL := `insert into users2(name,password) values ($1,$2)`
	result, err := tx.ExecContext(ctx, SQL, user.Name, string(hashpassword))
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `select password from users2 where name= $1`
	rows, err := tx.QueryContext(ctx, SQL, user.Name)
	helper.PanicIfError(err)
	defer rows.Close()

	if user.Id == 0 {
		return user, errors.New("user not found")
	}

	userHash := domain.User{}
	err = rows.Scan(&userHash.Password)

	err = bcrypt.CompareHashAndPassword([]byte(userHash.Password), []byte(user.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `update users set name && password = $1 where id = $2`
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := `delete from users where id = $1`
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := `select id,name from users2 where id = $1`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		helper.PanicIfError(err)
		defer rows.Close()

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
