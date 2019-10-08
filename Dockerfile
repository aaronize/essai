FROM elf-base:latest

WORKDIR /app/workspace/essai/

COPY ./essai-api /app/workspace/essai/

# RUN

EXPOSE 13030 8080

ENV ESSAI_RUN_MOD "devel"

ENTRYPOINT ["/app/workspace/essai/essai-api"]

