package graph

import (
	"context"
	"log"

	"crapi.community/graphql.grpc/graph/model"
	"crapi.community/graphql.grpc/grpc/models"

	pb "crapi.community/graphql.grpc/grpc/community-api/src/main/proto/com/crapi/services/community/stub/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateCoupon(address string, cp model.Coupon) (bool, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return false, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	created, err := client.CreateCoupon(ctx, &pb.CreateCouponRequest{
		Coupon: &pb.Coupon{
			CouponCode: cp.CouponCode,
			Amount:     cp.Amount,
			CreatedAt:  timestamppb.New(cp.CreatedAt),
		}})

	if err != nil {
		log.Println("Failed creating a Coupon, %v", err)
		return false, err
	}
	return created.GetSuccess(), nil
}

func UpdateCoupon(address string, id string, coupon model.CouponInput) (bool, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return false, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	uc := models.PrepareUpdatedCoupon(coupon)

	updated, err := client.UpdateCoupon(ctx, &pb.UpdateCouponRequest{
		CouponCode: id,
		UpdatedCoupon: &pb.Coupon{
			CouponCode: uc.CouponCode,
			Amount:     uc.Amount,
			CreatedAt:  timestamppb.New(uc.CreatedAt),
		},
	})

	if err != nil {
		log.Println("Coupon Updation Failed, %v", err)
		return false, err
	}
	return updated.GetSuccess(), nil
}

func GetCoupon(address string, ids []string) *pb.GetCouponsResponse {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return &pb.GetCouponsResponse{}
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	GetCoupon, err := client.GetCoupons(ctx, &pb.GetCouponsRequest{
		CouponCodes: ids,
	})

	if err != nil {
		log.Println("Get Coupons failed, %v", err)
		return nil
	}
	return &pb.GetCouponsResponse{
		Coupons: GetCoupon.Coupons,
	}
}

func DeleteCoupon(address string, CouponCode []string) ([]string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	DeleteCoupon, err := client.DeleteCoupons(ctx, &pb.DeleteCouponsRequest{
		CouponCodes: CouponCode,
	})

	res := []string{}
	if err != nil {
		log.Println("DeleteCoupon failed")
	}
	for i := 0; i < len(DeleteCoupon.DeletedCoupons); i++ {
		res = append(res, DeleteCoupon.DeletedCoupons[i].CouponCode)
	}
	return res, nil
}
