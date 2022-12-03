FROM nginx:stable-alpine

COPY css.css /bin/www
COPY js.js /bin/www
COPY index.html /bin/www
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]