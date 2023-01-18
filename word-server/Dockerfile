FROM rust:latest as build-env
ENV PKG_CONFIG_ALLOW_CROSS=1

WORKDIR /app
COPY . /app
RUN cargo build --release


FROM gcr.io/distroless/cc
COPY --from=build-env /app/target/release/word-server /
COPY --from=build-env /app/dictionary.db /
COPY --from=build-env /app/index.html /
EXPOSE 8080
CMD ["./word-server"]
