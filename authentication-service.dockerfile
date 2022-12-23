
#1.
FROM alpine:latest

#2. Run command mkdir on the new small docker image
RUN mkdir /app

#3. copy
COPY authApp /app

#4. Run the command
CMD [ "/app/authApp" ]

