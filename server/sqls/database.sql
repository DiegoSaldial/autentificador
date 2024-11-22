
create table `usuarios`(
    `id` integer unsigned auto_increment not null primary key,
    `nombres` varchar(60) not null,
    `apellido1` varchar(30) not null,
    `apellido2` varchar(30),
    `documento` varchar(30),
    `celular` varchar(20),
    `correo` varchar(100),
    `sexo` char(1),
    `direccion` varchar(100),
    `estado` tinyint(1) not null default 1,
    `username` varchar(30) unique not null,
    `password` varchar(64) not null, -- hash
    `last_login` datetime,
    `oauth_id` varchar(80),
    `foto_url` varchar(80),
    `ubicacion` point,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    `fecha_update` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE now()
);

create table `roles`(
    `nombre` varchar(50) not null primary key,
    `descripcion` varchar(100),
    `jerarquia` tinyint(1) not null default 0,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00')
);

-- para permisos no hay crud, se agregan segun se crean las funcionaes 
-- en el propio codigo fuente 
create table `permisos`(
    -- en golang hay funciones(query,mutation), cada funcion representa un permiso, el nombre de esa funcion 
    -- es el valor de metodo, (ej createPersona)
    `metodo` varchar(50) not null primary key,
    -- el nombre es el mismo que metodo, pero para que un humano lo lea, (ej crear persona)
    `nombre` varchar(50) not null,
    `descripcion` varchar(200),
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00')
);

create table `rol_permiso`(
    `rol` varchar(50) not null,
    `metodo` varchar(50) not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`rol`) references `roles`(`nombre`),
    foreign key(`metodo`) references `permisos`(`metodo`),
    primary key(`rol`,`metodo`)
);

create table `rol_usuario`(
    `rol` varchar(50) not null,
    `usuario_id` integer unsigned not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`rol`) references `roles`(`nombre`),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    primary key(`rol`,`usuario_id`)
);

create table `usuario_permiso`(
    `usuario_id` integer unsigned not null,
    `metodo` varchar(50) not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`metodo`) references `permisos`(`metodo`),
    primary key(`usuario_id`,`metodo`)
);

create table `menus`(
    `id` tinyint unsigned auto_increment not null primary key,
    `label` varchar(40) not null unique,
    `path` varchar(40) not null,
    `icon` varchar(40) not null,
    `color` varchar(40) not null,
    `grupo` tinyint(1) unsigned not null default 1,
    `orden` tinyint(1) unsigned not null default 1
);

create table `menus_usuario`(
    `id` tinyint unsigned auto_increment not null primary key,
    `usuario_id` integer unsigned not null,
    `menu_id` tinyint unsigned not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`menu_id`) references `menus`(`id`)
);

create table `rol_menus`(
    `id` tinyint unsigned auto_increment not null primary key,
    `rol` varchar(50) not null,
    `menu_id` tinyint unsigned not null,
    `fecha_registro` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`rol`) references `roles`(`nombre`),
    foreign key(`menu_id`) references `menus`(`id`)
);













-- indice para optimizar la busqueda 
CREATE INDEX idx_username ON usuarios (username);

-- VALORES POR DEFECTO
INSERT INTO `usuarios` (`nombres`, `apellido1`, `username`, `password`)
VALUES
    ('User', 'Admin', 'admin', SHA2('admin', 256));

INSERT INTO `roles` (`nombre`, `descripcion`, `jerarquia`)
VALUES
    ('Administrador', 'Tiene acceso total al sistema.', 0),
    ('Invitado', 'Acceso parcial, generalmente se autentifica con cuenta de google', 1);

INSERT INTO `permisos` (`metodo`, `nombre`, `descripcion`)
VALUES
    ('createUsuario', 'Crear Usuario', 'Permite crear un nuevo usuarios en el sistema.'),
    ('updateUsuario', 'Actualizar Usuario', 'Permite actualizar los datos de un usuarios en el sistema.'),
    ('updateUsuarioPerfil', 'Actualizar Perfil del Usuario', 'Permite actualizar los datos de un usuarios en el sistema.'),
    ('createRol', 'Crear Rol', 'Permite crear un nuevo rol en el sistema.'),
    ('updateRol', 'Actualizar Rol', 'Permite actualizar los datos de un rol en el sistema.'),
    ('roles', 'Listar roles', 'Listar los roles en el sistema.'),
    ('permisos', 'Listar permisos', 'Listar los permisos en el sistema.'),
    ('usuarios', 'Listar usuarios', 'Listar los usuarios en el sistema.'),
    ('usuarioById', 'Listar usuario por id', 'Listar los datos de un usuario en el sistema.'),
    ('rolById', 'Listar rol por id', 'Listar los datos de un rol en el sistema.'),
    ('asignarMenusUsuario', 'Asignar', 'Asignar permisos al usuario en el sistema.'),
    ('menus', 'menus', 'Listar menus en el sistema.'),
    ('menus_by_usuario', 'Listar menus por usuario', 'Menus de usuario en el sistema.');

insert into `rol_permiso`(`rol`,`metodo`)
values 
    ('Administrador','createUsuario'),
    ('Administrador','updateUsuario'),
    ('Administrador','updateUsuarioPerfil'),
    ('Administrador','createRol'),
    ('Administrador','updateRol'),
    ('Administrador','roles'),
    ('Administrador','permisos'),
    ('Administrador','usuarios'),
    ('Administrador','usuarioById'),
    ('Administrador','rolById'),
    ('Administrador','asignarMenusUsuario'),
    ('Administrador','menus'),
    ('Administrador','menus_by_usuario'),
    ('Invitado','updateUsuarioPerfil'),
    ('Invitado','usuarios'),
    ('Invitado','permisos'), 
    ('Invitado','roles'), 
    ('Invitado','usuarioById'), 
    ('Invitado','menus'),
    ('Invitado','menus_by_usuario');

insert into `rol_usuario`(`rol`,`usuario_id`)
values 
    ('Administrador',1);




insert into `menus`(`id`,`label`,`path`,`icon`,`grupo`,`color`,`orden`) values 
(1,'Usuarios','/usuarios','group',1,'primary',1),
(2,'Roles','/roles','local_movies',1,'primary',2);


insert into `menus_usuario`(usuario_id,menu_id) values 
(1,1),
(1,2);

insert into `rol_menus`(rol,menu_id) values 
('Invitado',1);





