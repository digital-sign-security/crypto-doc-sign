package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	client clients.Client
	logger *logrus.Logger
}

func NewUserRepository(logger *logrus.Logger, client clients.Client) *UserRepository {
	return &UserRepository{logger: logger, client: client}
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]*domains.UserWithKey, error) {
	q := `
		SELECT 
			u.id, u.username, k.public_key 
		FROM 
			public.user as u 
		LEFT JOIN 
			public.keys as k 
		ON 
			k.user_id=u.id
		WHERE 
			k.is_alive=true
	`
	u.logger.Infof("SQL Query: %s", formatQuery(q))

	rows, err := u.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]*domains.UserWithKey, 0)

	for rows.Next() {
		var user domains.UserWithKey

		err = rows.Scan(&user.ID, &user.Username, &user.PublicKey)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetUser(ctx context.Context, username, password string) (*domains.User, error) {
	q := `
		SELECT id, username, email, password FROM public.user WHERE username = $1 and password = $2
	`
	u.logger.Infof("SQL Query: %s", formatQuery(q))

	var user domains.User
	err := u.client.QueryRow(ctx, q, username, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		u.logger.Errorf("SQL Query: %s - err: %s", formatQuery(q), err)
		return &domains.User{}, fmt.Errorf("cannot get query: %w", err)
	}

	return &user, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user *domains.User) (*domains.User, error) {
	q := `
		INSERT INTO public.user 
		    (username, email, password) 
		VALUES 
			($1, $2, $3) 
		RETURNING id
	`
	u.logger.Infof("SQL Query: %s", formatQuery(q))
	if err := u.client.QueryRow(ctx, q, user.Username, user.Email, user.Username).Scan(&user.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			u.logger.Error(newErr)
			return nil, newErr
		}
		return nil, fmt.Errorf("cannot insert query: %w", err)
	}

	return user, nil
}

func (u *UserRepository) PatchUser(ctx context.Context) (*domains.User, error) {
	// TODO: make patch
	return nil, nil
}
