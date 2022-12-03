FROM nginx:stable-alpine

COPY css.css /var/www/html/
COPY js.js /var/www/html/
COPY index.html /var/www/html/
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]