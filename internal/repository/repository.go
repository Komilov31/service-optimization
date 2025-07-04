package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Komilov31/l0/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	pgxpool *pgxpool.Pool
	logger  *zap.Logger
}

func New(pgxpool *pgxpool.Pool, logger *zap.Logger) *Repository {
	return &Repository{pgxpool: pgxpool, logger: logger}
}

func (r *Repository) GetOrderById(ctx context.Context, uid uuid.UUID) (model.Order, error) {
	query := `
    SELECT
        order_uid,
        track_number,
        entry,
        locale,
        internal_signature,
        customer_id,
        delivery_service,
        shardkey,
        sm_id,
        date_created,
        oof_shard
	FROM orders
	WHERE order_uid = $1`

	tx, err := r.pgxpool.Begin(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return model.Order{}, fmt.Errorf("could not start transaction to db: %w", err)
	}
	defer tx.Rollback(ctx)

	var order model.Order
	err = r.pgxpool.QueryRow(ctx, query, uid).Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	)
	if err != nil {
		r.logger.Error(err.Error())
		return model.Order{}, fmt.Errorf("could not get order info from db: %w", err)
	}

	delivery, err := r.getOrderDelivery(ctx, uid)
	if err != nil {
		r.logger.Error(err.Error())
		return model.Order{}, fmt.Errorf("could not get delivery info from db: %w", err)
	}

	payment, err := r.getOrderPayment(ctx, uid)
	if err != nil {
		return model.Order{}, fmt.Errorf("could not get payment info from db: %w", err)
	}

	items, err := r.getOrderItems(ctx, uid)
	if err != nil {
		r.logger.Error(err.Error())
		return model.Order{}, fmt.Errorf("could not get items info from db: %w", err)
	}

	order.Delivery = delivery
	order.Items = items
	order.Payment = payment

	if err := tx.Commit(ctx); err != nil {
		r.logger.Error(err.Error())
		return model.Order{}, fmt.Errorf("could not commit transcation: %w", err)
	}

	r.logger.Info(
		"succesfully got order from db",
		zap.Time("time", time.Now()),
	)

	return order, nil
}

func (r *Repository) GetLastOrders(ctx context.Context) ([]model.Order, error) {
	rowNum := 100
	query := fmt.Sprintf(`
    SELECT
        order_uid,
        track_number,
        entry,
        locale,
        internal_signature,
        customer_id,
        delivery_service,
        shardkey,
        sm_id,
        date_created,
        oof_shard
	FROM orders
	LIMIT %d`, rowNum)

	tx, err := r.pgxpool.Begin(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, fmt.Errorf("could not start transaction to db: %w", err)
	}
	defer tx.Rollback(ctx)

	var orders []model.Order
	rows, err := r.pgxpool.Query(ctx, query)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, fmt.Errorf("could not get orders from db: %w", err)
	}

	for rows.Next() {
		var order model.Order

		err := rows.Scan(
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			r.logger.Error(err.Error())
			return nil, fmt.Errorf("could not scan field to order: %w", err)
		}

		orders = append(orders, order)
	}

	for i := 0; i < len(orders); i++ {
		orderUid := orders[i].OrderUid

		delivery, err := r.getOrderDelivery(ctx, orderUid)
		if err != nil {
			r.logger.Error(err.Error())
			return nil, fmt.Errorf("could not get delivery info from db: %w", err)
		}

		payment, err := r.getOrderPayment(ctx, orderUid)
		if err != nil {
			r.logger.Error(err.Error())
			return nil, fmt.Errorf("could not get payment info from db: %w", err)
		}

		items, err := r.getOrderItems(ctx, orderUid)
		if err != nil {
			r.logger.Error(err.Error())
			return nil, fmt.Errorf("could not get items info from db: %w", err)
		}

		orders[i].Delivery = delivery
		orders[i].Items = items
		orders[i].Payment = payment
	}

	if err := tx.Commit(ctx); err != nil {
		r.logger.Error(err.Error())
		return nil, fmt.Errorf("could not commit transcation: %w", err)
	}

	r.logger.Info(
		"successfully got last orders from db",
		zap.Time("time", time.Now()),
	)

	return orders, nil
}

func (r *Repository) CreateOrder(ctx context.Context, order model.Order) error {
	tx, err := r.pgxpool.Begin(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not start transaction to db: %w", err)
	}
	defer tx.Rollback(ctx)

	err = r.insertDeliveryToDb(ctx, order.Delivery)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not insert delivery info to db: %w", err)
	}

	err = r.insertPaymentToDb(ctx, order.Payment)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not insert payment info to db: %w", err)
	}

	err = r.insertOrderToDb(ctx, order)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not insert order info to db: %w", err)
	}

	err = r.insertItemsToDb(ctx, order)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not insert items info to db: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("could not commit transcation: %w", err)
	}

	r.logger.Info(
		"sucessfully created order in db",
		zap.Time("time", time.Now()),
		zap.String("uid", order.OrderUid.String()),
	)

	return nil
}
