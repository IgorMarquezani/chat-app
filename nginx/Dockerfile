FROM igorxaviermarquezani/chat-app-react:latest AS react

FROM nginx:latest

COPY --from=react /app/build /usr/share/nginx/html

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

CMD [ "nginx", "-g", "daemon off;" ]
