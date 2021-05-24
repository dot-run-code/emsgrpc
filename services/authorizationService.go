package services

import (
	"context"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JwtClaims struct {
	AuthorizedParty string `json:"azp,omitempty"`
	jwt.StandardClaims
}

var jwtKey []byte

type AuthoizationService struct {
	clientID           string
	clientSecret       string
	validAudiences     string
	issuer             string
	serviceRoleMapping map[string][]string
}

func NewAuthoizationService(clientId, clientSecret, validAudiences, issuer string, serviceRoleMapping map[string][]string) *AuthoizationService {
	jwtKey = []byte(clientSecret)
	return &AuthoizationService{
		clientID:           clientId,
		clientSecret:       clientSecret,
		validAudiences:     validAudiences,
		issuer:             issuer,
		serviceRoleMapping: serviceRoleMapping,
	}
}
func (a *AuthoizationService) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
func (a *AuthoizationService) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		_, err := validateContext(ss.Context())
		if err != nil {
			return err
		}
		return handler(srv, ss)
	}
}

func validateContext(ctx context.Context) (*JwtClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "No authentocation token provided")
	}
	values := md.Get("authorization")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "No authentocation token provided")
	}
	var claims = &JwtClaims{}

	token, _ := jwt.ParseWithClaims(values[0], claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "unexpected header alg")
		}
		return jwtKey, nil
	})
	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}
	return claims, nil

}
