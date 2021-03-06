// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/favorite.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

// Folder represents a favorite top-level folder in kbfs.
// This type is likely to change significantly as all the various parts are
// connected and tested.
type Folder struct {
	Name            string `codec:"name" json:"name"`
	Private         bool   `codec:"private" json:"private"`
	NotificationsOn bool   `codec:"notificationsOn" json:"notificationsOn"`
	Created         bool   `codec:"created" json:"created"`
}

type FavoritesResult struct {
	FavoriteFolders []Folder `codec:"favoriteFolders" json:"favoriteFolders"`
	IgnoredFolders  []Folder `codec:"ignoredFolders" json:"ignoredFolders"`
	NewFolders      []Folder `codec:"newFolders" json:"newFolders"`
}

type FavoriteAddArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type FavoriteIgnoreArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type GetFavoritesArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type FavoriteInterface interface {
	// Adds a folder to a user's list of favorite folders.
	FavoriteAdd(context.Context, FavoriteAddArg) error
	// Removes a folder from a user's list of favorite folders.
	FavoriteIgnore(context.Context, FavoriteIgnoreArg) error
	// Returns all of a user's favorite folders.
	GetFavorites(context.Context, int) (FavoritesResult, error)
}

func FavoriteProtocol(i FavoriteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.favorite",
		Methods: map[string]rpc.ServeHandlerDescription{
			"favoriteAdd": {
				MakeArg: func() interface{} {
					ret := make([]FavoriteAddArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FavoriteAddArg)
					if !ok {
						err = rpc.NewTypeError((*[]FavoriteAddArg)(nil), args)
						return
					}
					err = i.FavoriteAdd(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"favoriteIgnore": {
				MakeArg: func() interface{} {
					ret := make([]FavoriteIgnoreArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FavoriteIgnoreArg)
					if !ok {
						err = rpc.NewTypeError((*[]FavoriteIgnoreArg)(nil), args)
						return
					}
					err = i.FavoriteIgnore(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getFavorites": {
				MakeArg: func() interface{} {
					ret := make([]GetFavoritesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetFavoritesArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetFavoritesArg)(nil), args)
						return
					}
					ret, err = i.GetFavorites(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type FavoriteClient struct {
	Cli rpc.GenericClient
}

// Adds a folder to a user's list of favorite folders.
func (c FavoriteClient) FavoriteAdd(ctx context.Context, __arg FavoriteAddArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteAdd", []interface{}{__arg}, nil)
	return
}

// Removes a folder from a user's list of favorite folders.
func (c FavoriteClient) FavoriteIgnore(ctx context.Context, __arg FavoriteIgnoreArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteIgnore", []interface{}{__arg}, nil)
	return
}

// Returns all of a user's favorite folders.
func (c FavoriteClient) GetFavorites(ctx context.Context, sessionID int) (res FavoritesResult, err error) {
	__arg := GetFavoritesArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.favorite.getFavorites", []interface{}{__arg}, &res)
	return
}
