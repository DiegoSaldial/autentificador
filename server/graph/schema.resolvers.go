package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"auth/database/auth/menu"
	"auth/database/auth/permisos"
	"auth/database/auth/roles"
	"auth/database/auth/usuarios"
	"auth/database/auth/xauth"
	"auth/database/auth/xlogin"
	"auth/graph/model"
	"context"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.NewLogin) (*model.ResponseLogin, error) {
	return xlogin.Login(r.DB, input)
}

// Refreshtoken is the resolver for the refreshtoken field.
func (r *mutationResolver) Refreshtoken(ctx context.Context, token string, refreshToken string) (string, error) {
	return xlogin.RefreshToken(token, refreshToken)
}

// CreateUsuario is the resolver for the createUsuario field.
func (r *mutationResolver) CreateUsuario(ctx context.Context, input model.NewUsuario) (*model.Usuario, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "createUsuario")
	if err != nil {
		return nil, err
	}
	return usuarios.Crear(r.DB, input)
}

// UpdateUsuario is the resolver for the updateUsuario field.
func (r *mutationResolver) UpdateUsuario(ctx context.Context, input model.UpdateUsuario) (*model.Usuario, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "updateUsuario")
	if err != nil {
		return nil, err
	}
	return usuarios.Actualizar(r.DB, input)
}

// CreateRol is the resolver for the createRol field.
func (r *mutationResolver) CreateRol(ctx context.Context, input model.NewRol) (*model.ResponseRolCreate, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "createRol")
	if err != nil {
		return nil, err
	}
	return roles.Crear(r.DB, input)
}

// UpdateRol is the resolver for the updateRol field.
func (r *mutationResolver) UpdateRol(ctx context.Context, input model.NewRol) (*model.ResponseRolCreate, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "updateRol")
	if err != nil {
		return nil, err
	}
	return roles.Actualizar(r.DB, input)
}

// AsignarMenusUsuario is the resolver for the asignarMenusUsuario field.
func (r *mutationResolver) AsignarMenusUsuario(ctx context.Context, input model.AsignarMenusUsuario) (string, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "asignarMenusUsuario")
	if err != nil {
		return "", err
	}
	return menu.AsignarMenusUsuario(r.DB, input)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context, input model.InputMe) (*model.ResponseMe, error) {
	tok, err := xauth.CtxValue(ctx, r.DB, "")
	if err != nil {
		return nil, err
	}
	userid := tok.Usuario.ID
	return usuarios.GetMe(r.DB, input, userid)
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context, showPermisos bool) ([]*model.ResponseRolCreate, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "roles")
	if err != nil {
		return nil, err
	}
	return roles.GetRoles(r.DB, showPermisos)
}

// Permisos is the resolver for the permisos field.
func (r *queryResolver) Permisos(ctx context.Context) ([]*model.Permiso, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "permisos")
	if err != nil {
		return nil, err
	}
	return permisos.GetPermisos(r.DB)
}

// Usuarios is the resolver for the usuarios field.
func (r *queryResolver) Usuarios(ctx context.Context, query model.QueryUsuarios) ([]*model.Usuario, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "usuarios")
	if err != nil {
		return nil, err
	}
	return usuarios.GetUsuarios(r.DB, query)
}

// UsuarioByID is the resolver for the usuarioById field.
func (r *queryResolver) UsuarioByID(ctx context.Context, input model.GetUser) (*model.ResponseMe, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "usuarioById")
	if err != nil {
		return nil, err
	}
	return usuarios.GetById2(r.DB, input)
}

// RolByID is the resolver for the rolById field.
func (r *queryResolver) RolByID(ctx context.Context, rol string) (*model.ResponseRolCreate, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "rolById")
	if err != nil {
		return nil, err
	}
	return roles.GetRolById2(r.DB, rol, true)
}

// Menus is the resolver for the menus field.
func (r *queryResolver) Menus(ctx context.Context) ([]*model.Menus, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "menus")
	if err != nil {
		return nil, err
	}
	return menu.Listar(r.DB)
}

// MenusByUsuario is the resolver for the menus_by_usuario field.
func (r *queryResolver) MenusByUsuario(ctx context.Context, id string) ([]*model.Menus, error) {
	_, err := xauth.CtxValue(ctx, r.DB, "menus_by_usuario")
	if err != nil {
		return nil, err
	}
	return menu.ListarAsignados(r.DB, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
