package repository

import (
	"context"
	"errors"
	"example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	DB *mongo.Database
}

func NewPayment(db *mongo.Database) Payment {
	return Payment{DB: db}
}

func (p Payment) FindNotPaid() ([]model.Invoice, error) {
	invoices := []model.Invoice{}
	coll := p.DB.Collection("invoices")
	filter := bson.M{"status": "CREATED"}

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return invoices, nil
	}

	err = cursor.All(context.TODO(), &invoices)
	if err != nil {
		return invoices, nil
	}

	return invoices, nil
}

func (p Payment) FindByIdInvoice(id primitive.ObjectID) (model.Invoice, error) {
	invoice := model.Invoice{}
	coll := p.DB.Collection("invoices")

	filter := bson.M{"_id": id}
	err := coll.FindOne(context.Background(), filter).Decode(&invoice)

	if err != nil {
		return invoice, err
	}

	return invoice, nil
}

func (p Payment) PaidInvoice(invId primitive.ObjectID) error {
	coll := p.DB.Collection("invoices")

	update := bson.M{
		"$set": bson.M{"status": "PAID"},
	}

	_, err := coll.UpdateByID(context.Background(), invId, update)

	if err != nil {
		return errors.New("failed to update")
	}

	return nil
}

func (p Payment) PaymentFindByInvoiceId(invId primitive.ObjectID) (model.Payment, error) {
	payment := model.Payment{}
	coll := p.DB.Collection("payment")

	filter := bson.M{"invoice_id": invId}
	err := coll.FindOne(context.Background(), filter).Decode(&payment)

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (p Payment) CreatePayment(invoice_id primitive.ObjectID, amount int) (model.Payment, error) {
	newPayment := model.Payment{
		InvoiceId: invoice_id,
		Amount:    amount,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	coll := p.DB.Collection("payment")
	result, err := coll.InsertOne(context.Background(), newPayment)
	if err != nil {
		return newPayment, err
	}

	newPayment.ID = result.InsertedID.(primitive.ObjectID)
	return newPayment, nil
}
