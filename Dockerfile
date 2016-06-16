FROM scratch
ADD planckservice /
CMD ["-addr", ":8000"]
ENTRYPOINT ["/planckservice"]
