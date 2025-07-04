package repository

import (
	"context"

	"github.com/Komilov31/l0/internal/model"
	"github.com/google/uuid"
)

func (r *Repository) getOrderDelivery(ctx context.Context, orderUid uuid.UUID) (model.Delivery, error) {
	row := r.pgxpool.QueryRow(
		ctx,
		`SELECT * FROM deliveries 
		WHERE delivery_uid = (
		SELECT delivery_uid FROM orders
		WHERE order_uid = $1)`,
		orderUid,
	)

	var delivery model.Delivery
	err := row.Scan(
		&delivery.DeliveryUid,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	)
	if err != nil {
		return model.Delivery{}, err
	}

	return delivery, nil
}

func (r *Repository) getOrderItems(ctx context.Context, orderUid uuid.UUID) ([]model.Item, error) {
	rows, err := r.pgxpool.Query(
		ctx,
		`SELECT i.*
	FROM items i
	JOIN order_items oi ON oi.item_uid = i.item_uid
	WHERE oi.order_uid = $1;`,
		orderUid)
	if err != nil {
		return nil, err
	}

	var items []model.Item
	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.ItemUid,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) getOrderPayment(ctx context.Context, orderUid uuid.UUID) (model.Payment, error) {
	row := r.pgxpool.QueryRow(
		ctx,
		`SELECT * FROM payments
		WHERE payment_uid = (
		SELECT payment_uid FROM orders
		WHERE order_uid = $1)`,
		orderUid,
	)

	var payment model.Payment
	err := row.Scan(
		&payment.PaymentUid,
		&payment.Transaction,
		&payment.RequestId,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func (r *Repository) insertDeliveryToDb(ctx context.Context, delivery model.Delivery) error {
	query := `INSERT INTO deliveries (delivery_uid, name, phone, zip, city, address, region, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.pgxpool.Exec(
		ctx,
		query,
		delivery.DeliveryUid,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) insertPaymentToDb(ctx context.Context, payment model.Payment) error {
	query := `INSERT INTO payments (payment_uid, transaction, request_id, currency, 
	provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.pgxpool.Exec(
		ctx,
		query,
		payment.PaymentUid,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) insertItemsToDb(ctx context.Context, order model.Order) error {
	query := `INSERT INTO items (item_uid, chrt_id, track_number, price, 
	rid, name, sale, size, total_price, nm_id, brand, status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	items := order.Items
	for _, item := range items {
		_, err := r.pgxpool.Exec(
			ctx,
			query,
			item.ItemUid,
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
		)
		if err != nil {
			return err
		}

		_, err = r.pgxpool.Exec(
			ctx,
			"INSERT INTO order_items(order_uid, item_uid) VALUES ($1, $2)",
			order.OrderUid,
			item.ItemUid,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) insertOrderToDb(ctx context.Context, order model.Order) error {
	query := `INSERT INTO orders (order_uid, track_number, entry, delivery_uid, 
	payment_uid, locale, internal_signature, customer_id, delivery_service, 
	shardkey, sm_id, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := r.pgxpool.Exec(
		ctx,
		query,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery.DeliveryUid,
		order.Payment.PaymentUid,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.ShardKey,
		order.SmId,
		order.OofShard,
	)
	if err != nil {
		return err
	}

	return nil
}
