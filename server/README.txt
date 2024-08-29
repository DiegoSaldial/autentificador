==========================================================
NOTAS DE MANTENIBILIDAD Y AMPLIACION
==========================================================

El alcance de este sistema es **usuarios**, **roles** y **permisos**
Generando token y refreshToken de autentificacion 


Al crear una nueva funcionalidad el nombre de la funcion 
debe ir como un registro en la tabla permisos de la db
esa funcionalidad se llama desde el archivo 'schema.resolvers.go'

Por ejemplo para el permiso **'usuarioById'** existe un query con el mismo
nombre **'usuarioById(input:GetUser!): ResponseMe!'**



