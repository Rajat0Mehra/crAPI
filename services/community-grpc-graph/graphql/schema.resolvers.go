package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

/***
This file contains resolver functions for all Query and Mutations of graphql.
*/

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"crapi.community/graphql.grpc/graphql/auth"
	"crapi.community/graphql.grpc/graphql/config"
	"crapi.community/graphql.grpc/graphql/model"
	"crapi.community/graphql.grpc/grpc/models"
	pb "crapi.community/graphql.grpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	post := models.PrepareNewPost(model.Post{
		Title:   input.Title,
		Content: input.Content,
	})

	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return &model.Post{}, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)
	p := models.Convert_Graph_pb_post(post)
	_, err = client.CreatePost(ctx, &pb.CreatePostRequest{
		Post: p,
	})

	if err != nil {
		log.Println("Failed creating a post, %v", err)
		return &model.Post{}, err
	}
	ret := models.Convert_pb_Graph_post(p)
	return &ret, nil
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.PostInput) (bool, error) {
	post := models.PrepareUpdatePost(input)
	res, err := UpdatePost(os.Getenv("GRPC-SERVICE"), id, post)

	if err != nil {
		log.Println("Error while updating post....%v", err)
		return false, err
	}
	return res, nil
}

// DeletePost is the resolver for the DeletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postsID []string) ([]string, error) {
	res, err := DeletePost(os.Getenv("GRPC-SERVICE"), postsID)

	if err != nil {
		log.Println("Error while deleting post ..... %v", err)
		return []string{}, err
	}
	return res, err
}

// CreateCoupon is the resolver for the CreateCoupon field.
func (r *mutationResolver) CreateCoupon(ctx context.Context, input model.CouponInput) (*model.Coupon, error) {
	coupon := models.PrepareNewCoupon(input)

	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return &model.Coupon{}, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	_, err = client.CreateCoupon(ctx, &pb.CreateCouponRequest{
		Coupon: &pb.Coupon{
			CouponCode: coupon.CouponCode,
			Amount:     coupon.Amount,
			CreatedAt:  timestamppb.New(coupon.CreatedAt),
		}})

	if err != nil {
		log.Println("Failed creating a Coupon, %v", err)
		return &model.Coupon{}, err
	}
	return &coupon, nil
}

// UpdateCoupon is the resolver for the UpdateCoupon field.
func (r *mutationResolver) UpdateCoupon(ctx context.Context, id string, input model.CouponInput) (bool, error) {
	res, err := UpdateCoupon(os.Getenv("GRPC-SERVICE"), id, input)

	if err != nil {
		log.Println("Error while Updating Coupon by graphQL ..... %v", err)
		return false, err
	}
	return res, nil
}

// DeleteCoupon is the resolver for the DeleteCoupon field.
func (r *mutationResolver) DeleteCoupon(ctx context.Context, coupons []string) ([]string, error) {
	res, err := DeleteCoupon(os.Getenv("GRPC-SERVICE"), coupons)

	if err != nil {
		log.Println("Error while deleting post ..... %v", err)
		return []string{}, err
	}
	return res, err
}

// AddComment is the resolver for the AddComment field.
func (r *mutationResolver) AddComment(ctx context.Context, input *model.CommentInput) (*model.Post, error) {
	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return &model.Post{}, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	user := models.PrepareUser()
	_, err = client.CreateComment(ctx, &pb.CreateCommentRequest{
		Comment: &pb.Comment{
			Id:        input.ID,
			Content:   input.Content,
			CreatedAt: timestamppb.New(time.Now()),
			Author: &pb.User{
				Nickname:  user.Nickname,
				Email:     user.Email,
				VehicleId: user.Vehicleid,
				Picurl:    user.ProfilePicURL,
				CreatedAt: timestamppb.New(user.CreatedAt),
			},
		},
	})

	if err != nil {
		log.Println("Can not add comment to post")
		return &model.Post{}, err
	}
	retPost, err := client.GetPosts(ctx, &pb.GetPostsRequest{
		Ids: []string{input.ID},
	})
	if err != nil {
		log.Println("Failed to fetch the post after adding comment, %v", err)
		return &model.Post{}, err
	}

	p_model := models.Convert_pb_Graph_post(retPost.GetPosts()[0])
	return &p_model, nil
}

// GetPosts is the resolver for the GetPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, ids []string) ([]*model.Post, error) {
	// lis := []string{"oBvaQhfFqVWVdDBgAhaER6", "MnwSHsuN9okaPEiSerDuA3"}
	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	GetPosts, err := client.GetPosts(ctx, &pb.GetPostsRequest{
		Ids: ids,
	})

	if err != nil {
		println("GetPosts() failed .... recieved no response from grpc server ... , %v", err)
		return nil, err
	}

	retPost := []*model.Post{}
	for i := 0; i < len(GetPosts.GetPosts()); i++ {
		convertedP := models.Convert_pb_Graph_post(GetPosts.GetPosts()[i])
		retPost = append(retPost, &convertedP)
	}

	return retPost, nil
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context, limit int) ([]*model.Post, error) {
	// In this function, we are bypassing the connection with grpc_client and directly communicating with grpc_server
	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewCommunityServiceClient(conn)

	ctx = context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	GetPosts, err := client.GetAllPosts(ctx, &pb.GetAllPostRequest{
		NoOfPosts: int32(limit),
	})
	if err != nil {
		log.Println("GetPost failed ... could not get posts from grpc")
		return nil, err
	}
	retPost := []*model.Post{}
	for i := 0; i < len(GetPosts.GetPosts()); i++ {
		convertedP := models.Convert_pb_Graph_post(GetPosts.GetPosts()[i])
		retPost = append(retPost, &convertedP)
	}
	return retPost, nil
}

// GetCoupons is the resolver for the GetCoupons field.
func (r *queryResolver) GetCoupons(ctx context.Context, codes []string) ([]*model.Coupon, error) {
	req := ctx.Value("request").(*http.Request)
	_, err := auth.ExtractTokenID(req, server.DB)
	if err != nil {
		log.Fatalln("Unauthorized")
		return nil, err
	}

	Coupons := GetCoupon(os.Getenv("GRPC-SERVICE"), codes)

	ret := []*model.Coupon{}
	for i := 0; i < len(Coupons.Coupons); i++ {
		c := model.Coupon{
			CouponCode: Coupons.Coupons[i].CouponCode,
			Amount:     Coupons.Coupons[i].Amount,
			CreatedAt:  models.Convert_to_Time(Coupons.Coupons[i].CreatedAt),
		}
		ret = append(ret, &c)
	}
	return ret, nil
}

// ValidateCoupon is the resolver for the ValidateCoupon field.
func (r *queryResolver) ValidateCoupon(ctx context.Context, code string) (*model.Coupon, error) {
	if code == "" {
		return nil, errors.New("Coupon Code empty")
	}

	conn, err := grpc.Dial(os.Getenv("GRPC-SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Println("Cannot connect to server, %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)
	ctx = context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	log.Println("All connections have been made .... validating coupon now")
	coupons, err := client.GetCoupons(ctx, &pb.GetCouponsRequest{
		CouponCodes: []string{code},
	})
	if err != nil {
		log.Println("Coupon does not exist")
		return nil, err
	}
	return &model.Coupon{
		CouponCode: coupons.GetCoupons()[0].GetCouponCode(),
		Amount:     coupons.GetCoupons()[0].GetAmount(),
		CreatedAt:  models.Convert_to_Time(coupons.GetCoupons()[0].GetCreatedAt()),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var server = config.Server{}
