package repository_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	mocks "mockerydb/mocks/repository"
	"mockerydb/models"
)

func TestRepository_UpdatePayment(t *testing.T) {
	type args struct {
		id      string
		payment models.Payment
	}
	tests := []struct {
		name    string
		args    args
		want    models.Payment
		wantErr bool
	}{
		{
			name: "success_update_payment",
			args: args{
				id: "0000-0000-0000-000",
				payment: models.Payment{
					Status: "PAID",
				},
			},
			want: models.Payment{
				Status: "PAID",
			},
			wantErr: false,
		},
		{
			name: "failed_update_payment",
			args: args{
				id: "0000-0000-0000-000",
				payment: models.Payment{
					Status: "PAID",
				},
			},
			want:    models.Payment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.IPayementRepository{}
			if !tt.wantErr {
				repo.On("UpdatePayment", tt.args.id, tt.args.payment).Return(tt.want, nil)
			} else {
				repo.On("UpdatePayment", tt.args.id, tt.args.payment).Return(tt.want, errors.New("update error"))
			}

			got, err := repo.UpdatePayment(tt.args.id, tt.args.payment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.UpdatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_DeletePayment(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success_delete_payment",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "failed_delete_payment",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.IPayementRepository{}
			if !tt.wantErr {
				repo.On("DeletePayment", tt.args.id).Return(tt.want, nil)
			} else {
				repo.On("DeletePayment", tt.args.id).Return(tt.want, errors.New("delete is failed"))
			}
			got, err := repo.DeletePayment(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.DeletePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.DeletePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SelectPaymentWIthId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Payment
		wantErr bool
	}{
		{
			name: "success_select_payment_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want: models.Payment{
				Status:      "PENDING",
				PaymentCode: "payment-code-001",
				Created:     time.Now(),
				Updated:     time.Now(),
			},
			wantErr: false,
		},
		{
			name: "failed_select_payment_by_id",
			args: args{
				id: "0000-0000-0000-000",
			},
			want:    models.Payment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.IPayementRepository{}
			if !tt.wantErr {
				repo.On("SelectPaymentWithId", tt.args.id).Return(tt.want, nil)
			} else {
				repo.On("SelectPaymentWithId", tt.args.id).Return(tt.want, errors.New("Failed to get payment by id"))
			}
			got, err := repo.SelectPaymentWithId(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.SelectPaymentWithId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.SelectPaymentWithId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreatePayment(t *testing.T) {
	type args struct {
		payment models.Payment
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success_create_payment",
			args: args{
				models.Payment{
					PaymentCode: "payment-code-011",
					Status:      "PENDING",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "failed_create_payment",
			args: args{
				models.Payment{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.IPayementRepository{}
			if !tt.wantErr {
				repo.On("CreatePayment", tt.args.payment).Return(tt.want, nil)
			} else {
				repo.On("CreatePayment", tt.args.payment).Return(tt.want, errors.New("Failed to create payment"))
			}
			got, err := repo.CreatePayment(tt.args.payment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.CreatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
