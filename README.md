# Playground

- Falta implementar sistema de logging, las opciones son implementar una libreria como zap de uber o crear nuestra propia solución (consulta a GPT)
- Implementar test para refactorizar con seguridad
- Implementar cookie/auth middleware para evitar hacer uso de id para acceder a los recursos:
    Ideas para las cookies: generarlas a mano con un uuid metodo new ya genera uno de 128 bits, crear un map de string[string] llamado db
    y usar la agrupación de rutas para limitar el acceso a las rutas