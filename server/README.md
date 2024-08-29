# Notas de Mantenibilidad y Ampliación

El alcance de este sistema incluye **usuarios**, **roles** y **permisos**.

Generando el **token** y **refreshToken** de autenticación.

Al crear una nueva funcionalidad, el nombre de la función debe registrarse en la tabla de permisos de la base de datos. Esa funcionalidad se invoca desde el archivo `schema.resolvers.go`.

Por ejemplo, para el permiso **`usuarioById`**, existe un query con el mismo nombre: **`usuarioById(input:GetUser!): ResponseMe!`**.
