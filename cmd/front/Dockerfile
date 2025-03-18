# Etapa 1: Construcción de la aplicación Angular
FROM node:20 AS build

WORKDIR /app

# Copiar archivos de dependencias
COPY package.json package-lock.json ./

# Instalar dependencias
RUN npm ci

# Copiar el código fuente de Angular
COPY . .

# Construir la aplicación en modo producción
RUN npm run build -- --configuration=production

# Etapa 2: Servir con Nginx
FROM nginx:stable-alpine

# Copiar archivos de configuración de Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copiar la aplicación construida desde la etapa anterior
COPY --from=build /app/dist/access-gate /usr/share/nginx/html

# Exponer el puerto 80
EXPOSE 80

# Iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
